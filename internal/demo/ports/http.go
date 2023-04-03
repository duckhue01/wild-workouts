package ports

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/duckhue01/wild-workouts/internal/common/server/httperr"
	"github.com/duckhue01/wild-workouts/internal/demo/app"
	"github.com/duckhue01/wild-workouts/internal/demo/app/query"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}

func (h HttpServer) ListAllDemos(w http.ResponseWriter, r *http.Request) {
	resp, err := h.app.Queries.AllDemos.Handle(r.Context(), query.AllDemos{})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
	}

	demos := make([]Demo, 0, len(resp))

	for i := 0; i < len(resp); i++ {
		demos = append(demos, Demo{Name: resp[i].Name})
	}

	render.Respond(w, r, demos)
}
