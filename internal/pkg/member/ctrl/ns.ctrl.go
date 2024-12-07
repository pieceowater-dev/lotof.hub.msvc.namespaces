package ctrl

import "app/internal/pkg/member/svc"

type MemberController struct {
	memberService *svc.MemberService
	//pb.UnimplementedMemberServiceServer
}

func NewMemberController(service *svc.MemberService) *MemberController {
	return &MemberController{
		memberService: service,
	}
}
