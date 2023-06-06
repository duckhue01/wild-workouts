package app

import "github.com/tribefintech/microservices/internal/demo/app/query"

type Application struct {
	// Commands Commands
	Queries Queries
}

// type Commands struct {
// 	MakeDemo command.UpdateDemohandler
// }

type Queries struct {
	AllDemos query.ListCurrentUserDemosHandler
}
