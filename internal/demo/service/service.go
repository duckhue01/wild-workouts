package service

import (
	"context"
	"database/sql"

	"github.com/duckhue01/wild-workouts/internal/common/metrics"
	"github.com/duckhue01/wild-workouts/internal/demo/adapters"
	"github.com/duckhue01/wild-workouts/internal/demo/app"
	"github.com/duckhue01/wild-workouts/internal/demo/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	hourRepository := adapters.NewPostgresRepository(sql.OpenDB(nil))
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.NoOp{}

	return app.Application{
		Queries: app.Queries{
			AllDemo: query.NewAllTrainingsHandler(hourRepository, logger, metricsClient),
		},
	}
}
