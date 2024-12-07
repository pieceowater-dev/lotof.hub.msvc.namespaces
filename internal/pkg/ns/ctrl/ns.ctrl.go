package ctrl

import "app/internal/pkg/ns/svc"

type NSController struct {
	nsService *svc.NSService
	//pb.UnimplementedNamespaceServiceServer
}

func NewNSController(service *svc.NSService) *NSController {
	return &NSController{
		nsService: service,
	}
}
