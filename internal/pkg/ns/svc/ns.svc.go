package svc

import gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"

type NSService struct {
	db gossiper.Database
}

func NewNSService(db gossiper.Database) *NSService {
	return &NSService{db: db}
}
