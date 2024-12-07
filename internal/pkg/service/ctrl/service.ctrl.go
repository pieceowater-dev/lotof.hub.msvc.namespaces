package ctrl

import "app/internal/pkg/service/svc"

type ServiceController struct {
	serviceService *svc.ServiceService
	//pb.UnimplementedServiceServiceServer
}

func NewServiceController(service *svc.ServiceService) *ServiceController {
	return &ServiceController{
		serviceService: service,
	}
}
