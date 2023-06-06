package service

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/tribefintech/microservices/internal/common/metrics"

	"github.com/tribefintech/microservices/internal/demo/adapter"
	"github.com/tribefintech/microservices/internal/demo/app"
	"github.com/tribefintech/microservices/internal/demo/app/query"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"host" `
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`
	Env     string `mapstructure:"env"`
	Port    int    `mapstructure:"port"`
	Cognito struct {
		Region string `mapstructure:"region"`
		PoolId string `mapstructure:"poolId"`
	} `mapstructure:"cognito"`
}

type Secret struct {
	Database struct {
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`
}

func NewApplication(ctx context.Context, sec *Secret, conf *Config) app.Application {
	logger := logrus.NewEntry(logrus.StandardLogger())
	logger.Info("config and secret is loaded")

	db, err := sql.Open("postgres", "postgres://postgres:postgres@postgres:5432/tribe?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	hourRepository := adapter.NewPostgres(db)
	metricsClient := metrics.NoOp{}

	return app.Application{
		Queries: app.Queries{
			AllDemos: query.NewAllDemosHandler(hourRepository, logger, metricsClient),
		},
	}
}
