package pkg

import (
	"app/internal/pkg/member"
	//pb "app/internal/core/grpc/generated"
	"app/internal/pkg/ns"
	"github.com/gin-gonic/gin"
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

func (r *Router) Init(grpcServer *grpc.Server) {
	// Register gRPC services
	//pb.RegisterNamespaceServiceServer(grpcServer, r.todoModule.Controller)
}

// InitGRPC initializes gRPC routes
func (r *Router) InitGRPC(grpcServer *grpc.Server) {
	// Register gRPC services
	//pb.RegisterNamespaceServiceServer(grpcServer, r.todoModule.Controller)
}

// InitREST initializes REST routes using Gin
func (r *Router) InitREST(router *gin.Engine) {
	//api := router.Group("/api")
	{
		// Register GIN routes
		//api.GET("/todos", r.todoModule.Controller.ListREST)
	}
}
