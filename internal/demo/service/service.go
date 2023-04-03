package service

import (
	"context"
	"fmt"

	"github.com/duckhue01/wild-workouts/internal/common/logs"
	"github.com/duckhue01/wild-workouts/internal/common/metrics"
	"github.com/duckhue01/wild-workouts/internal/demo/adapter"
	"github.com/duckhue01/wild-workouts/internal/demo/app"
	"github.com/duckhue01/wild-workouts/internal/demo/app/query"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"host" `
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`
	Env string `mapstructure:"env"`
}

type Secret struct {
	Database struct {
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`
}

var config = new(Config)
var secret = new(Secret)

func init() {
	err := readSecret()
	if err != nil {
		panic(err)
	}

	err = readConfig()
	if err != nil {
		panic(err)
	}

	logs.Init(config.Env)

}

func NewApplication(ctx context.Context) app.Application {
	logger := logrus.NewEntry(logrus.StandardLogger())
	logger.Info("config and secret is loaded")

	hourRepository := adapter.NewMemory()
	metricsClient := metrics.NoOp{}

	return app.Application{
		Queries: app.Queries{
			AllDemos: query.NewAllDemosHandler(hourRepository, logger, metricsClient),
		},
	}
}

func readConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config file: %w", err)
	}

	if err := viper.UnmarshalExact(config, func(dc *mapstructure.DecoderConfig) {
		dc.ErrorUnset = true
	}); err != nil {
		return fmt.Errorf("unmarshal config: %w", err)
	}
	return nil
}

func readSecret() error {
	viper.SetConfigName("secret")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read secret file: %w", err)
	}

	if err := viper.UnmarshalExact(secret, func(dc *mapstructure.DecoderConfig) {
		dc.ErrorUnset = true
	}); err != nil {
		return fmt.Errorf("unmarshal secret: %w", err)
	}
	return nil
}
