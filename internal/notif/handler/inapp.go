package handler

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/tribefintech/microservices/internal/notif/domain"
	"github.com/tribefintech/microservices/internal/notif/domain/notifs"
	"github.com/tribefintech/microservices/internal/notif/ports"
)

func (h *Handler) inAppHandler(notifcs <-chan *notifs.Notif) {
	logger := h.logger.WithFields(logrus.Fields{
		"channel": domain.InAppChannel,
	})

	for n := range notifcs {
		logger.Info("received a notification")

		br, err := json.Marshal(h.buildWSPayload(n, logger))
		if err != nil {
			logger.Errorf("error marshal message: %s", err.Error())
		}

		err = h.wsp.Push(n.ReceiverId.String(), br)
		if err != nil {
			logger.Errorf("can not push notif: %s", err.Error())
		}

		logger.Info("pushed to websocket channel")
	}
}

func (h *Handler) buildWSPayload(n *notifs.Notif, logger *logrus.Entry) interface{} {
	var payload interface{}
	switch ndata := n.Data.(type) {
	case *notifs.SocialMessageCreatedNotifData:
		payload = &ports.SocialMessageCreated{
			Event: ports.EventSocialMessageCreated,
			Msg:   ndata.Msg,
			Sender: struct {
				FirstName string "json:\"first_name\""
				Id        string "json:\"id\""
				ImageUrl  string "json:\"image_url\""
				LastName  string "json:\"last_name\""
			}{
				FirstName: ndata.Sender.FirstName,
				Id:        ndata.Sender.Id.String(),
				ImageUrl:  ndata.Sender.ImageURL,
				LastName:  ndata.Sender.LastName,
			},
			ChannelId: ndata.ChannelId,
			CreateAt:  ndata.CreateAt,
		}
	case *notifs.EndorsementRequestCreatedNotifData:
		payload = &ports.EndorsementRequestCreated{
			Event: ports.EventSocialMessageCreated,
			Sender: struct {
				FirstName string "json:\"first_name\""
				Id        string "json:\"id\""
				ImageUrl  string "json:\"image_url\""
				LastName  string "json:\"last_name\""
			}{
				FirstName: ndata.Sender.FirstName,
				Id:        ndata.Sender.Id.String(),
				ImageUrl:  ndata.Sender.ImageURL,
				LastName:  ndata.Sender.LastName,
			},
			CreateAt:      ndata.CreateAt,
			EndorsementId: ndata.EndorsementId,
		}
	case *notifs.EducationCourseProgressRemindData:
		payload = &ports.EducationCourseProgressRemind{
			Course: &struct {
				Id    int    "json:\"id\""
				Title string "json:\"title\""
			}{
				Id:    ndata.Course.Id,
				Title: ndata.Course.Title,
			},
			CreateAt: ndata.CreateAt,
			Event:    ports.EventEducationCourseProgressRemind,
		}
	default:
		logger.Warning("can not cast the type of notification data")
	}

	return payload
}
