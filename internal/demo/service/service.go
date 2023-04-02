package service

import (
	"context"

	"github.com/duckhue01/wild-workouts/internal/common/metrics"
	"github.com/duckhue01/wild-workouts/internal/demo/adapter"
	"github.com/duckhue01/wild-workouts/internal/demo/app"
	"github.com/duckhue01/wild-workouts/internal/demo/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	hourRepository := adapter.NewMemory()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NoOp{}

	return app.Application{
		Queries: app.Queries{
			AllDemos: query.NewAllDemosHandler(hourRepository, logger, metricsClient),
		},
	}
}
