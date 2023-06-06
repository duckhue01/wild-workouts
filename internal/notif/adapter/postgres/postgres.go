package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/tabbed/pqtype"

	"github.com/tribefintech/microservices/internal/notif/domain"
	"github.com/tribefintech/microservices/internal/notif/domain/notifs"
	"github.com/tribefintech/microservices/internal/notif/postgres/sqlc"
)

const (
	migrationDir = "file:///internal/notif/postgres/migration"
)

type Postgres struct {
	db     *sql.DB
	sql    *sqlc.Queries
	logger *logrus.Entry
}

func NewPostgres(db *sql.DB, logger *logrus.Entry) *Postgres {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationDir, "postgres", driver)
	if err != nil {
		panic(err)
	}

	err = m.Up()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}

	logger.Info("database migrated")

	return &Postgres{
		db:  db,
		sql: sqlc.New(db),
		logger: logger.WithFields(logrus.Fields{
			"adapter": "postgres",
		}),
	}
}

func (p *Postgres) InsertNotification(ctx context.Context, n *notifs.Notif) error {
	prams, err := p.marshalNotif(*n)
	if err != nil {
		return err
	}

	err = p.sql.InsertNotification(ctx, prams)

	return err
}

func (p *Postgres) ListNotification(ctx context.Context) (*[]notifs.Notif, error) {
	res, err := p.sql.ListUserNotifications(ctx, sqlc.ListUserNotificationsParams{
		ReceiverID: uuid.NullUUID{},
		ID:         [16]byte{},
		Limit:      0,
	})

	notis := make([]notifs.Notif, 0)

	for _, v := range res {
		notis = append(notis, p.unmarshalNotif(v))
	}

	return &notis, err

}

func (p *Postgres) unmarshalNotif(n sqlc.NotifNotification) notifs.Notif {
	return notifs.Notif{
		Seen:       n.IsSeen.Bool,
		Event:      domain.EventCode(n.Event.String),
		ReceiverId: n.ReceiverID.UUID,
		SenderId:   n.SenderID,
		Data:       n.Data,
	}
}

func (p *Postgres) marshalNotif(n notifs.Notif) (sqlc.InsertNotificationParams, error) {
	bdata, err := n.ByteData()

	return sqlc.InsertNotificationParams{
		ID: uuid.New(),
		ReceiverID: uuid.NullUUID{
			UUID:  n.ReceiverId,
			Valid: true,
		},
		Data: pqtype.NullRawMessage{
			RawMessage: bdata,
			Valid:      true,
		},
		IsSeen: sql.NullBool{
			Bool:  false,
			Valid: true,
		},
		Event: sql.NullString{
			String: string(n.Event),
			Valid:  true,
		},
		SenderID: uuid.NullUUID{
			UUID:  n.SenderId.UUID,
			Valid: n.SenderId.Valid,
		},
	}, err
}
