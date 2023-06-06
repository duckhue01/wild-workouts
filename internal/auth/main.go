package main

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/go-chi/chi/v5"
	"github.com/tribefintech/microservices/internal/common/logs"
	"github.com/tribefintech/microservices/internal/common/server"

	cmsvc "github.com/tribefintech/microservices/internal/common/service"
)

type Config struct {
	Env     string `mapstructure:"env"`
	Port    int    `mapstructure:"port"`
	Cognito struct {
		Region   string `mapstructure:"region"`
		ClientId string `mapstructure:"clientId"`
	} `mapstructure:"cognito"`
}

var conf = new(Config)

func init() {
	var err error

	conf, err = cmsvc.ReadConfig[Config](".", "config", "yaml")
	if err != nil {
		panic(err)
	}

	logs.Init(conf.Env)
}

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(conf.Cognito.Region),
	}))

	// ctx := context.Background()
	cognito := newCognito(conf.Cognito.ClientId, cognitoidentityprovider.New(sess))
	server.Run(server.Conf{
		CreateHandler: func(router chi.Router) http.Handler {
			return HandlerFromMux(newHTTPServer(cognito), router)
		},
		Port: conf.Port,
	})
}
