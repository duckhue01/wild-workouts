package handler

import (
	"github.com/sirupsen/logrus"
	"github.com/tribefintech/microservices/internal/notif/domain"
	"github.com/tribefintech/microservices/internal/notif/domain/notifs"
)

type PushPayload struct {
	Title   string
	Message string
	UserId  string
}

func (h *Handler) pushHandler(notifcs <-chan *notifs.Notif) {
	logger := h.logger.WithFields(logrus.Fields{
		"channel": domain.PushChannel,
	})

	for n := range notifcs {
		logger.Info("received a notification")
		h.p.Push(n.Event, n.ReceiverId.String())
	}
}
