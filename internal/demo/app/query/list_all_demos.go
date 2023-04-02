package query

import (
	"context"

	"github.com/duckhue01/wild-workouts/internal/common/decorator"
	"github.com/sirupsen/logrus"
)

type AllDemos struct{}

type AllDemosHandler decorator.QueryHandler[AllDemos, []*Demo]

type allDemosHandler struct {
	readModel AllDemosReadModel
}

func NewAllDemosHandler(
	readModel AllDemosReadModel,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) AllDemosHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[AllDemos, []*Demo](
		allDemosHandler{readModel: readModel},
		logger,
		metricsClient,
	)
}

type AllDemosReadModel interface {
	ListAllDemos(ctx context.Context) ([]*Demo, error)
}

func (h allDemosHandler) Handle(ctx context.Context, _ AllDemos) (tr []*Demo, err error) {
	return h.readModel.ListAllDemos(ctx)
}
