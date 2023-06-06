package ports

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tribefintech/microservices/internal/common/auth"
	"github.com/tribefintech/microservices/internal/common/cmerr"
	"golang.org/x/net/websocket"
)

const (
	ErrPushToNilConnection = "push-to-nil-connection"
)

const (
	PingFrame = "ping"
	PongFrame = "pong"
)

type WS struct {
	logger *logrus.Entry
	conns  map[string]*websocket.Conn
}

func NewWS(logger *logrus.Entry) *WS {
	return &WS{
		logger: logger.WithFields(logrus.Fields{
			"port": "web-socket",
		}),
		conns: make(map[string]*websocket.Conn),
	}
}

func (w *WS) SubscribeNotification(c *websocket.Conn) {
	defer c.Close()
	user, err := auth.UserFromCtx(c.Request().Context())

	logger := w.logger.WithFields(logrus.Fields{
		"userId": user.UUID,
	})
	if err != nil {
		logger.Errorf("non-identity user: %s", err.Error())
	}
	w.conns[user.UUID] = c
	logger.Info("subscribed to notification service")

	w.ping(c, &user, logger)
}

func (w *WS) CheckUserIsActive(userId string) bool {
	if _, ok := w.conns[userId]; ok {
		return true
	}

	return false
}

func (w *WS) Push(userId string, data []byte) error {
	conn, ok := w.conns[userId]
	if !ok {
		return cmerr.New("push to nil connection", ErrPushToNilConnection, cmerr.TypUnexpected)
	}

	_, err := conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (w *WS) ping(c *websocket.Conn, user *auth.User, logger *logrus.Entry) {
	defer func() {
		delete(w.conns, user.UUID)
		err := c.Close()
		if err != nil {
			logger.Errorf("can not close connection: %s", err)
			return
		}

		logger.Info("connection closed")
	}()

	for {
		err := websocket.Message.Send(c, PingFrame)
		if err != nil {
			logger.Warnf("can not ping: %s", err)
			return
		}

		var pong string
		err = websocket.Message.Receive(c, &pong)
		if err != nil {
			logger.Warnf("can not receive msg: %s", err)
			return
		}

		if pong != PongFrame {
			logger.Warnf("not receive pong")
			return
		}

		time.Sleep(5 * time.Second)
	}
}
