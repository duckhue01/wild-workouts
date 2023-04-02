package adapter

import (
	"context"
	"database/sql"

	"github.com/duckhue01/wild-workouts/internal/demo/app/query"
	"github.com/duckhue01/wild-workouts/internal/demo/postgres/sqlc"
)

type Postgres struct {
	db  *sql.DB
	sql *sqlc.Queries
}

func NewPostgres(db *sql.DB) *Postgres {

	return &Postgres{
		db:  db,
		sql: sqlc.New(db),
	}
}

func (a *Postgres) ListAllDemos(ctx context.Context) ([]*query.Demo, error) {
	res, err := a.sql.ListAllDemos(ctx)
	if err != nil {
		return nil, err
	}

	demos := make([]*query.Demo, 0, len(res))

	for i := 0; i < len(res); i++ {
		demos = append(demos, &query.Demo{
			ID:   res[i].ID,
			Name: res[i].Name,
		})
	}

	return demos, err
}
