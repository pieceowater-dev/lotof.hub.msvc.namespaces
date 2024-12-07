package ns

import (
	"app/internal/core/cfg"
	"app/internal/pkg/ns/ctrl"
	"app/internal/pkg/ns/svc"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type Module struct {
	Controller *ctrl.NSController
}

func New() *Module {
	database, err := gossiper.NewDB(
		gossiper.PostgresDB,
		cfg.Inst().PostgresDatabaseDSN,
		false,
		[]any{},
	)
	if err != nil {
		log.Fatalf("Failed to create database instance: %v", err)
	}

	return &Module{
		Controller: ctrl.NewNSController(
			svc.NewNSService(database),
		),
	}
}
