package logs

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

const (
	LocalEnv = "local"
)

func Init(env string) {
	SetFormatter(logrus.StandardLogger(), env)

	logrus.SetLevel(logrus.DebugLevel)
}

func SetFormatter(logger *logrus.Logger, env string) {
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	})

	if env == LocalEnv {
		logger.SetFormatter(&prefixed.TextFormatter{
			ForceFormatting: true,
		})
	}
}
