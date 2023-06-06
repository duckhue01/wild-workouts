package ports

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/tribefintech/microservices/internal/common/server/httperr"

	"github.com/tribefintech/microservices/internal/demo/app"
	"github.com/tribefintech/microservices/internal/demo/app/query"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}

func (h HttpServer) ListCurrentUserDemos(w http.ResponseWriter, r *http.Request, params ListCurrentUserDemosParams) {

	resp, err := h.app.Queries.AllDemos.Handle(r.Context(), query.ListCurrentUserDemosQuery{
		WantError: params.Error,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	demos := make([]Demo, 0, len(resp))

	for i := 0; i < len(resp); i++ {
		demos = append(demos, Demo{Name: resp[i].Name})
	}

	render.Respond(w, r, demos)
}

func (h HttpServer) CreateCurrentUserDemo(w http.ResponseWriter, r *http.Request, params CreateCurrentUserDemoParams) {

}
