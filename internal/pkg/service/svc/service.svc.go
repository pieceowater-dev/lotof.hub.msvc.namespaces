package svc

import gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"

type ServiceService struct {
	db gossiper.Database
}

func NewServiceService(db gossiper.Database) *ServiceService {
	return &ServiceService{db: db}
}
