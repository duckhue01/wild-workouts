package query

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/tribefintech/microservices/internal/common/cmerr"
	"github.com/tribefintech/microservices/internal/common/decorator"
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
		return nil, cmerr.New(
			err.Error(),
			cmerr.InternalServerError,
			cmerr.TypUnexpected,
		)
	}

	return demos, err
}
