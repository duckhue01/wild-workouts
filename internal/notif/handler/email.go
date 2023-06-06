package handler

import (
	"github.com/tribefintech/microservices/internal/notif/domain/notifs"
)

func (h *Handler) emailHandler(notifcs <-chan *notifs.Notif) {
	// logger := h.logger.WithFields(logrus.Fields{
	// 	"channel": domain.EmailChannel,
	// })

}
