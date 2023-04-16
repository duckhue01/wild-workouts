package query

import (
	"context"

	"github.com/duckhue01/wild-workouts/internal/common/cmerror"
	"github.com/duckhue01/wild-workouts/internal/common/decorator"
	"github.com/sirupsen/logrus"
)

type ListCurrentUserDemosQuery struct {
	WantError bool
}

type ListCurrentUserDemosHandler decorator.QueryHandler[ListCurrentUserDemosQuery, []*Demo]

type allDemosHandler struct {
	rm ListCurrentUserDemosRM
}

func NewAllDemosHandler(
	rm ListCurrentUserDemosRM,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) ListCurrentUserDemosHandler {
	if rm == nil {
		panic("nil read model")
	}

	return decorator.ApplyQueryDecorators[ListCurrentUserDemosQuery, []*Demo](
		allDemosHandler{rm: rm},
		logger,
		metricsClient,
	)
}

type ListCurrentUserDemosRM interface {
	ListAllDemos(context.Context, ListCurrentUserDemosQuery) ([]*Demo, error)
}

func (h allDemosHandler) Handle(ctx context.Context, q ListCurrentUserDemosQuery) (tr []*Demo, err error) {
	demos, err := h.rm.ListAllDemos(ctx, q)
	if err != nil {
		return nil, cmerror.New(
			err.Error(),
			cmerror.InternalServerError,
			cmerror.TypDomainError,
		)
	}

	return demos, err
}
