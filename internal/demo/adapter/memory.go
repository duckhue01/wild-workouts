package adapter

import (
	"context"

	"github.com/duckhue01/wild-workouts/internal/demo/app/query"
)

type Memory struct {
}

func NewMemory() *Memory {

	return &Memory{}
}

func (a *Memory) ListAllDemos(_ context.Context) ([]*query.Demo, error) {

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
