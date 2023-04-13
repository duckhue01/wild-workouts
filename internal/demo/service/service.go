package service

import (
	"context"

	"github.com/duckhue01/wild-workouts/internal/common/logs"
	"github.com/duckhue01/wild-workouts/internal/common/metrics"
	"github.com/duckhue01/wild-workouts/internal/common/service"
	"github.com/duckhue01/wild-workouts/internal/demo/adapter"
	"github.com/duckhue01/wild-workouts/internal/demo/app"
	"github.com/duckhue01/wild-workouts/internal/demo/app/query"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"host" `
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`
	Env string `mapstructure:"env"`
}

type Secret struct {
	Database struct {
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`
}

var config = new(Config)
var secret = new(Secret)

func init() {
	var err error

	secret, err = service.ReadSecret[Secret](".", "config", "yaml")
	if err != nil {
		panic(err)
	}

	config, err = service.ReadConfig[Config](".", "config", "yaml")
	if err != nil {
		panic(err)
	}

	logs.Init(config.Env)
}

func NewApplication(ctx context.Context) app.Application {
	logger := logrus.NewEntry(logrus.StandardLogger())
	logger.Info("config and secret is loaded")

	hourRepository := adapter.NewMemory()
	metricsClient := metrics.NoOp{}

	return app.Application{
		Queries: app.Queries{
			AllDemos: query.NewAllDemosHandler(hourRepository, logger, metricsClient),
		},
	}
}
