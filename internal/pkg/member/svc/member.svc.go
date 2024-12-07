package svc

import gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"

type MemberService struct {
	db gossiper.Database
}

func NewMemberService(db gossiper.Database) *MemberService {
	return &MemberService{db: db}
}
