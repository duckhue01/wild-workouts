package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/duckhue01/wild-workouts/internal/common/logs"
	"github.com/duckhue01/wild-workouts/internal/common/server"
	"github.com/duckhue01/wild-workouts/internal/demo/ports"
	"github.com/duckhue01/wild-workouts/internal/demo/service"
	"github.com/go-chi/chi/v5"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"host" `
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
	} `mapstructure:"database"`
}

type Secret struct {
	Database struct {
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`
}

var config = new(Config)
var secret = new(Secret)

func init() {
	readSecret()
	readConfig()

	logs.Init()
}

func main() {
	ctx := context.Background()

	app := service.NewApplication(ctx)
	server.Run(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(app), router)
	})

}

func readConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.UnmarshalExact(config, func(dc *mapstructure.DecoderConfig) {
		dc.ErrorUnset = true
	}); err != nil {
		panic(fmt.Errorf("fatal can not unmarshal config: %w", err))
	}
}

func readSecret() {
	viper.SetConfigName("secret")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error secret file: %w", err))
	}

	if err := viper.UnmarshalExact(secret, func(dc *mapstructure.DecoderConfig) {
		dc.ErrorUnset = true
	}); err != nil {
		panic(fmt.Errorf("fatal can not unmarshal secret: %w", err))
	}
}
