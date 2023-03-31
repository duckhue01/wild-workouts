package app

import "github.com/duckhue01/wild-workouts/internal/demo/app/query"

type Application struct {
	// Commands Commands
	Queries Queries
}

// type Commands struct {
// 	MakeDemo command.UpdateDemohandler
// }

type Queries struct {
	AllDemo query.AllDemosHandler
}
