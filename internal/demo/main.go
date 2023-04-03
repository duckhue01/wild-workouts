package main

import (
	"context"
	"net/http"

	"github.com/duckhue01/wild-workouts/internal/common/server"
	"github.com/duckhue01/wild-workouts/internal/demo/ports"
	"github.com/duckhue01/wild-workouts/internal/demo/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	ctx := context.Background()

	app := service.NewApplication(ctx)
	server.Run(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(app), router)
	}, 3000)
}
