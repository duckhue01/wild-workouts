package decorator

import (
	"context"
	"fmt"

	"github.com/duckhue01/wild-workouts/internal/common/cmerror"
	"github.com/sirupsen/logrus"
)

type commandLoggingDecorator[C any] struct {
	base   CommandHandler[C]
	logger *logrus.Entry
}

func (d commandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	handlerType := generateActionName(cmd)

	logger := d.logger.WithFields(logrus.Fields{
		"command":      handlerType,
		"command_body": fmt.Sprintf("%#v", cmd),
	})

	logger.Debug("executing command")
	defer func() {
		if err == nil {
			logger.Info("command executed successfully")
		} else {
			logger.WithError(err).Error("Failed to execute command")
		}
	}()

	return d.base.Handle(ctx, cmd)
}

type queryLoggingDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	logger *logrus.Entry
}

func (d queryLoggingDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	logger := d.logger.WithFields(logrus.Fields{
		"query":      generateActionName(cmd),
		"query_body": fmt.Sprintf("%#v", cmd),
	})

	logger.Debug("Executing query")
	defer func() {
		if err == nil {
			logger.Info("query executed successfully")
			return
		}

		if cmerr, ok := err.(cmerror.Error); ok {
			switch cmerr.ErrorType() {
			case cmerror.TypUnexpected:
				logger.WithError(err).Error("failed to execute query")
			default:
				logger.WithError(err).Warning("failed to execute query")
			}

			return
		}

		logger.WithError(err).Warning("failed to execute query")
	}()

	return d.base.Handle(ctx, cmd)
}
