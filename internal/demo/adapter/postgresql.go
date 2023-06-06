package adapter

import (
	"context"
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/tribefintech/microservices/internal/demo/app/query"
	"github.com/tribefintech/microservices/internal/demo/postgres/sqlc"
)

type Postgres struct {
	db  *sql.DB
	sql *sqlc.Queries
}

func NewPostgres(db *sql.DB) *Postgres {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:///internal/demo/postgres/migration", "postgres", driver)
	if err != nil {
		panic(err)
	}

	err = m.Up()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}

	return &Postgres{
		db:  db,
		sql: sqlc.New(db),
	}
}

func (a *Postgres) ListAllDemos(ctx context.Context, q query.ListCurrentUserDemosQuery) ([]*query.Demo, error) {
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
