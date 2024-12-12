package pkg

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/member"
	"app/internal/pkg/ns"
	"google.golang.org/grpc"
)

type Router struct {
	nsModule     *ns.Module
	memberModule *member.Module
}

func NewRouter() *Router {
	return &Router{
		nsModule:     ns.New(),
		memberModule: member.New(),
	}
}

// InitGRPC initializes gRPC routes
func (r *Router) InitGRPC(grpcServer *grpc.Server) {
	// Register gRPC services
	pb.RegisterNamespaceServiceServer(grpcServer, r.nsModule.Controller)
	pb.RegisterMemberServiceServer(grpcServer, r.memberModule.Controller)
}
