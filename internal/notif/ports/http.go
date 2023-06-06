package ports

import (
	"net/http"

	"github.com/go-chi/render"
)

type httpServer struct {
}

func NewHTTPServer() *httpServer {
	return &httpServer{}
}

func (h *httpServer) GetNotifHealthInformation(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, "OK")
}

func (h *httpServer) GetListUserNotifications(w http.ResponseWriter, r *http.Request, params GetListUserNotificationsParams) {

}

// dummy method: do not edit
func (h *httpServer) SubscribeNotification(w http.ResponseWriter, r *http.Request) {}
