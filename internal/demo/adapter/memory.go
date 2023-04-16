package adapter

import (
	"context"
	"errors"

	"github.com/duckhue01/wild-workouts/internal/demo/app/query"
)

type Memory struct {
}

func NewMemory() *Memory {

	return &Memory{}
}

func (a *Memory) ListAllDemos(_ context.Context, q query.ListCurrentUserDemosQuery) ([]*query.Demo, error) {

	if q.WantError {
		return nil, errors.New("da ta bey deo get duoc demos")

	}

	return []*query.Demo{
		{
			ID:   0,
			Name: "Damwon Kia",
		},
		{
			ID:   0,
			Name: "Damwon Kia",
		},
	}, nil

}
