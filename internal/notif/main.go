package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"

	"github.com/tribefintech/microservices/internal/common/auth/cognito"
	"github.com/tribefintech/microservices/internal/common/logs"
	"github.com/tribefintech/microservices/internal/common/server"
	cmsvc "github.com/tribefintech/microservices/internal/common/service"

	"github.com/tribefintech/microservices/internal/notif/adapter/nats"
	"github.com/tribefintech/microservices/internal/notif/adapter/onesignal"
	"github.com/tribefintech/microservices/internal/notif/adapter/postgres"
	"github.com/tribefintech/microservices/internal/notif/handler"
	"github.com/tribefintech/microservices/internal/notif/ports"
)

type Config struct {
	Env  string `mapstructure:"env"`
	Port int    `mapstructure:"port"`
	NATS struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"nats"`
	Cognito struct {
		Region string `mapstructure:"region"`
		PoolId string `mapstructure:"poolId"`
	} `mapstructure:"cognito"`
	Postgres struct {
		Host         string `mapstructure:"host" `
		Port         int    `mapstructure:"port"`
		Username     string `mapstructure:"username"`
		DatabaseName string `mapstructure:"databaseName"`
	} `mapstructure:"postgres"`
	OneSignal struct {
		AppId string `mapstructure:"appId" `
	} `mapstructure:"oneSignal"`
}

type Secret struct {
	Postgres struct {
		Password string `mapstructure:"password"`
	} `mapstructure:"postgres"`
	OneSignal struct {
		APIKey string `mapstructure:"apiKey" `
	} `mapstructure:"oneSignal"`
}

var conf = new(Config)
var sec = new(Secret)

func init() {
	var err error

	conf, err = cmsvc.ReadConfig[Config](".", "config", "yaml")
	if err != nil {
		panic(err)
	}

	sec, err = cmsvc.ReadSecret[Secret](".", "secret", "yaml")
	if err != nil {
		panic(err)
	}

	logs.Init(conf.Env)
}

func main() {
	logger := logrus.NewEntry(logrus.StandardLogger())
	logger.Info("config is loaded")

	parser := cognito.New(conf.Cognito.Region, conf.Cognito.PoolId)

	ws := ports.NewWS(logger)

	queue := nats.NewNATS(fmt.Sprintf("nats://%s:%d", conf.NATS.Host, conf.NATS.Port), logger)

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		conf.Postgres.Username,
		sec.Postgres.Password,
		conf.Postgres.Host,
		conf.Postgres.Port,
		conf.Postgres.DatabaseName,
	))
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	pg := postgres.NewPostgres(db, logger)

	push := onesignal.NewOneSignal(sec.OneSignal.APIKey, conf.OneSignal.AppId, logger)

	h := handler.NewHandler(queue, ws, logger, pg, push)

	// todo: run method meaning that the method will run until the process is killed => must be run in go routine => remove outermost goroutine inside this method
	h.Run()

	server.Run(server.Conf{
		CreateHandler: func(router chi.Router) http.Handler {
			router.Handle("/notif/ws", websocket.Handler(ws.SubscribeNotification))
			return ports.HandlerFromMux(ports.NewHTTPServer(), router)
		},
		Port:        conf.Port,
		TokenParser: parser,
		RouteWhiteList: []string{
			"/notif/health",
		},
	})
}
