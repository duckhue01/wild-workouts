package adapter

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/duckhue01/wild-workouts/internal/demo/app/query"
	"github.com/duckhue01/wild-workouts/internal/demo/postgres/sqlc"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Postgres struct {
	db  *sql.DB
	sql *sqlc.Queries
}

func NewPostgres(db *sql.DB) *Postgres {

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(fmt.Errorf("fatal can not create driver instance: %w", err))
	}

	m, err := migrate.NewWithDatabaseInstance("file:///../postgres/migration", "postgres", driver)
	if err != nil {
		panic(fmt.Errorf("fatal can not create migration instance: %w", err))
	}

	err = m.Up()
	if err != nil {
		panic(fmt.Errorf("fatal can not migrate database: %w", err))
	}

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
