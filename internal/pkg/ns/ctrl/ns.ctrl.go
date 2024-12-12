package ctrl

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/ns/ent"
	"app/internal/pkg/ns/svc"
	"context"
	"github.com/google/uuid"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type NSController struct {
	nsService *svc.NSService
	pb.UnimplementedNamespaceServiceServer
}

func NewNSController(service *svc.NSService) *NSController {
	return &NSController{
		nsService: service,
	}
}

func (N NSController) GetNamespaces(_ context.Context, request *pb.GetNamespacesRequest) (*pb.GetNamespacesResponse, error) {
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
	namespace, count, err := N.nsService.GetNamespaces(filter)
	if err != nil {
		return nil, err
	}

	paginatedNamespaces := make([]*pb.Namespace, count)
	for i, u := range *namespace {
		paginatedNamespaces[i] = &pb.Namespace{
			Id:          u.ID.String(),
			Title:       u.Title,
			Slug:        u.Slug,
			Description: u.Description,
			Owner:       u.Owner.String(),
		}
	}

	return &pb.GetNamespacesResponse{
		Namespaces: &pb.PaginatedNamespaceList{
			Rows: paginatedNamespaces,
			Info: &pb.PaginationInfo{Count: int32(count)},
		},
	}, nil
}

func (N NSController) GetNamespace(_ context.Context, request *pb.GetNamespaceRequest) (*pb.Namespace, error) {
	namespace, err := N.nsService.GetNamespace(request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Namespace{
		Id:          namespace.ID.String(),
		Title:       namespace.Title,
		Slug:        namespace.Slug,
		Description: namespace.Description,
		Owner:       namespace.Owner.String(),
	}, nil
}

func (N NSController) CreateNamespace(_ context.Context, request *pb.NamespaceRequest) (*pb.Namespace, error) {
	nsToCreate := ent.Namespace{
		Title:       request.Title,
		Slug:        request.Slug,
		Description: request.Description,
	}
	namespace, err := N.nsService.CreateNamespace(&nsToCreate)
	if err != nil {
		return nil, err
	}
	return &pb.Namespace{
		Id:          namespace.ID.String(),
		Title:       namespace.Title,
		Slug:        namespace.Slug,
		Description: namespace.Description,
		Owner:       namespace.Owner.String(),
	}, nil
}

func (N NSController) UpdateNamespace(_ context.Context, request *pb.UpdateNamespaceRequest) (*pb.Namespace, error) {
	namespace, err := N.nsService.UpdateNamespace(&ent.Namespace{
		ID:          uuid.MustParse(request.Id),
		Title:       request.Title,
		Slug:        request.Slug,
		Description: request.Description,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Namespace{
		Id:          namespace.ID.String(),
		Title:       namespace.Title,
		Slug:        namespace.Slug,
		Description: namespace.Description,
		Owner:       namespace.Owner.String(),
	}, nil
}
