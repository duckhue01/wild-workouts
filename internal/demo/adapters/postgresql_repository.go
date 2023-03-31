package adapters

import (
	"context"
	"database/sql"

	"github.com/duckhue01/wild-workouts/internal/demo/domain"
	"github.com/duckhue01/wild-workouts/internal/demo/postgres/sqlc"
)

type PostgresRepository struct {
	db  *sql.DB
	sql *sqlc.Queries
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {

	return &PostgresRepository{
		db:  db,
		sql: sqlc.New(db),
	}
}

func (a *PostgresRepository) GetDemo(ctx context.Context, id int64) (*domain.Demo, error) {
	res, err := a.sql.GetDemo(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.Demo{Name: res.Name}, err
}
