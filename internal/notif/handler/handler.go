package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/tribefintech/microservices/internal/notif/domain"
	"github.com/tribefintech/microservices/internal/notif/domain/notifs"
)

type (
	QueueAdapter interface {
		Run() <-chan domain.Event
	}

	SESAdapter interface{}

	PushAdapter interface {
		Push(ecode domain.EventCode, userId string)
	}

	Postgres interface {
		InsertNotification(ctx context.Context, n *notifs.Notif) error
	}
)

type (
	WebsocketPort interface {
		CheckUserIsActive(userId string) bool
		Push(userId string, data []byte) error
	}
)

type Handler struct {
	chans  map[domain.ChannelCode]chan *notifs.Notif
	q      QueueAdapter
	wsp    WebsocketPort
	logger *logrus.Entry
	db     Postgres
	p      PushAdapter
}

func NewHandler(q QueueAdapter, wsp WebsocketPort, logger *logrus.Entry, db Postgres, p PushAdapter) *Handler {
	return &Handler{
		chans: map[domain.ChannelCode]chan *notifs.Notif{
			domain.EmailChannel: make(chan *notifs.Notif),
			domain.PushChannel:  make(chan *notifs.Notif),
			domain.InAppChannel: make(chan *notifs.Notif),
		},
		q:      q,
		wsp:    wsp,
		logger: logger,
		db:     db,
		p:      p,
	}
}

func (h *Handler) Run() {
	source := h.q.Run()
	go h.fanOut(source)

	for k, v := range h.chans {
		switch k {
		case domain.EmailChannel:
			go h.emailHandler(v)
		case domain.PushChannel:
			go h.pushHandler(v)
		case domain.InAppChannel:
			go h.inAppHandler(v)
		}
	}
}

func (h *Handler) fanOut(source <-chan domain.Event) {

	for e := range source {
		switch ne := e.(type) {
		case *domain.SocialMessageCreated:
			h.handleSocialMessageCreated(ne)
		case *domain.EndorsementRequestCreated:
			h.handleEndorsementRequestCreated(ne)
		case *domain.EducationCourseProgressRemind:
			h.handleEducationCourseProgressRemind(ne)

		default:
			h.logger.WithFields(logrus.Fields{
				"event": e.Event(),
			}).Warn("event not supported")
			continue
		}

	}
}

func (h *Handler) handleSocialMessageCreated(e *domain.SocialMessageCreated) {
	for _, rid := range *e.ReceiverIds {
		notif := notifs.NewNotif(
			rid,
			&notifs.SocialMessageCreatedNotifData{
				Sender: &notifs.Sender{
					Id:        e.Sender.Id,
					FirstName: e.Sender.FirstName,
					LastName:  e.Sender.LastName,
					ImageURL:  e.Sender.ImageURL,
				},
				Msg: e.Msg,
			},
			h.wsp.CheckUserIsActive(rid.String()),
			uuid.NullUUID{
				UUID:  e.Sender.Id,
				Valid: false,
			},
			e.Event(),
		)

		err := h.db.InsertNotification(context.Background(), notif)
		if err != nil {
			h.logger.WithFields(logrus.Fields{
				"event": e.Event(),
			}).Errorf("error insert notification: %s", err.Error())
		}
		h.chans[notif.Channel()] <- notif
	}
}

func (h *Handler) handleEndorsementRequestCreated(e *domain.EndorsementRequestCreated) {
	notif := notifs.NewNotif(
		e.ReceiverId,
		&notifs.EndorsementRequestCreatedNotifData{
			Sender: &notifs.Sender{
				Id:        e.Sender.Id,
				FirstName: e.Sender.FirstName,
				LastName:  e.Sender.LastName,
				ImageURL:  e.Sender.ImageURL,
			},

			CreateAt:      e.CreatedAt,
			EndorsementId: e.EndorsementId,
		},
		h.wsp.CheckUserIsActive(e.ReceiverId.String()),
		uuid.NullUUID{
			UUID:  e.Sender.Id,
			Valid: true,
		},
		e.Event(),
	)

	err := h.db.InsertNotification(context.Background(), notif)
	if err != nil {
		h.logger.WithFields(logrus.Fields{
			"event": e.Event(),
		}).Errorf("error insert notification: %s", err.Error())
	}
	h.chans[notif.Channel()] <- notif

}

func (h *Handler) handleEducationCourseProgressRemind(e *domain.EducationCourseProgressRemind) {
	notif := notifs.NewNotif(
		e.ReceiverId,
		&notifs.EducationCourseProgressRemindData{
			Course: &domain.Course{
				Id:    e.Course.Id,
				Title: e.Course.Title,
			},
			CreateAt: e.CreatedAt,
		},
		h.wsp.CheckUserIsActive(e.ReceiverId.String()),
		uuid.NullUUID{
			Valid: false,
		},
		e.Event(),
	)

	notif.SetChannel(e.Channel())

	err := h.db.InsertNotification(context.Background(), notif)
	if err != nil {
		h.logger.WithFields(logrus.Fields{
			"event": e.Event(),
		}).Errorf("error insert notification: %s", err.Error())
	}

	h.chans[notif.Channel()] <- notif
}
