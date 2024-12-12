package ctrl

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/member/svc"
	"context"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type MemberController struct {
	memberService *svc.MemberService
	pb.UnimplementedMemberServiceServer
}

func NewMemberController(service *svc.MemberService) *MemberController {
	return &MemberController{
		memberService: service,
	}
}

func (m MemberController) GetMembers(_ context.Context, request *pb.GetMembersRequest) (*pb.GetMembersResponse, error) {
	filter := gossiper.NewFilter[string](
		request.GetSearch(),
		gossiper.NewSort[string](
			request.GetSort().GetField(),
			gossiper.SortDirection(request.GetSort().GetDirection()),
		),
		gossiper.NewPagination(
			int(request.GetPagination().GetPage()),
			int(request.GetPagination().GetLength()),
		),
	)
	members, count, err := m.memberService.GetMembers(filter)
	if err != nil {
		return nil, err
	}

	paginatedMembers := make([]*pb.Member, count)
	for i, m := range *members {
		paginatedMembers[i] = &pb.Member{
			Id:     m.ID.String(),
			UserId: m.UserID.String(),
		}
	}

	return &pb.GetMembersResponse{
		Members: &pb.PaginatedMemberList{
			Rows: paginatedMembers,
			Info: &pb.PaginationInfo{Count: int32(count)},
		},
	}, nil
}

func (m MemberController) GetMember(_ context.Context, request *pb.GetMemberRequest) (*pb.GetMemberResponse, error) {
	membership, err := m.memberService.GetMember(request.MembershipId)
	if err != nil {
		return nil, err
	}
	return &pb.GetMemberResponse{Member: &pb.Member{
		Id:     membership.ID.String(),
		UserId: membership.UserID.String(),
	}}, nil
}

func (m MemberController) AddMemberToNamespace(_ context.Context, request *pb.MemberToNamespaceRequest) (*pb.OK, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemberController) RemoveMemberFromNamespace(_ context.Context, request *pb.MemberToNamespaceRequest) (*pb.OK, error) {
	//TODO implement me
	panic("implement me")
}
