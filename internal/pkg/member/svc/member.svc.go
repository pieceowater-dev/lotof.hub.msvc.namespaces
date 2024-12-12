package svc

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/member/ent"
	"fmt"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type MemberService struct {
	db gossiper.Database
}

func NewMemberService(db gossiper.Database) *MemberService {
	return &MemberService{db: db}
}

func (s MemberService) GetMembers(filter gossiper.Filter[string]) (*[]ent.Member, int64, error) {
	var members []ent.Member
	var count int64

	query := s.db.GetDB().Model(&ent.Member{})

	// Apply search filters
	if filter.Search != "" {
		//todo: implement search by ns
		fmt.Println(filter.Search, "search by ns not implemented yet")
	}

	// Count total records
	if err := query.Count(&count).Error; err != nil {
		return &[]ent.Member{}, 0, fmt.Errorf("failed to count: %w", err)
	}

	// Apply pagination
	query = query.Offset((filter.Pagination.Page - 1) * filter.Pagination.Length).Limit(filter.Pagination.Length)

	// Apply sorting dynamically
	if field := filter.Sort.Field; field != "" && gossiper.IsFieldValid(&ent.Member{}, field) {
		query = query.Order(fmt.Sprintf("%s %s", gossiper.ToSnakeCase(field), filter.Sort.Direction))
	}

	if err := query.Find(&members).Error; err != nil {
		return &[]ent.Member{}, 0, fmt.Errorf("failed to query: %w", err)
	}

	return &members, count, nil
}

func (s MemberService) GetMember(id string) (*ent.Member, error) {
	var member *ent.Member

	if err := s.db.GetDB().
		Model(&ent.Member{}).
		Where("id = ?", id).
		First(&member).Error; err != nil {
		return nil, err
	}

	return member, nil
}

func (s MemberService) AddMemberToNamespace() (*pb.OK, error) {
	//TODO implement me
	panic("implement me")
}

func (s MemberService) RemoveMemberFromNamespace() (*pb.OK, error) {
	//TODO implement me
	panic("implement me")
}
