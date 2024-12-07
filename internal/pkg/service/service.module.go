package service

import (
	"app/internal/core/cfg"
	"app/internal/pkg/service/ctrl"
	"app/internal/pkg/service/svc"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type Module struct {
	Controller *ctrl.ServiceController
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
		Controller: ctrl.NewServiceController(
			svc.NewServiceService(database),
		),
	}
}
