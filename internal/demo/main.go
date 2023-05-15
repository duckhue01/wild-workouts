package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/duckhue01/wild-workouts/internal/common/auth/cognito"
	"github.com/duckhue01/wild-workouts/internal/common/logs"
	"github.com/duckhue01/wild-workouts/internal/common/server"
	cmsvc "github.com/duckhue01/wild-workouts/internal/common/service"

	"github.com/duckhue01/wild-workouts/internal/demo/ports"
	"github.com/duckhue01/wild-workouts/internal/demo/service"
)

var conf = new(service.Config)
var sec = new(service.Secret)

func init() {
	var err error

	sec, err = cmsvc.ReadSecret[service.Secret](".", "secret", "yaml")
	if err != nil {
		panic(err)
	}

	conf, err = cmsvc.ReadConfig[service.Config](".", "config", "yaml")
	if err != nil {
		panic(err)
	}

	logs.Init(conf.Env)
}

func main() {
	ctx := context.Background()

	parser := cognito.New(conf.Cognito.Region, conf.Cognito.PoolId)

	app := service.NewApplication(ctx, sec, conf)
	server.Run(server.Conf{
		CreateHandler: func(router chi.Router) http.Handler {
			return ports.HandlerFromMux(ports.NewHttpServer(app), router)
		},
		Port:   conf.Port,
		Parser: parser,
	})
}
