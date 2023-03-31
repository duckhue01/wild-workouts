package ports

import (
	"net/http"

	"github.com/duckhue01/wild-workouts/internal/demo/app"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}

func (h HttpServer) GetAllDemos(w http.ResponseWriter, r *http.Request) {

}
