package routing

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"context"
	// "github.com/bav-demo/internal/controllers/society"

	"google.golang.org/grpc"

	pb "github.com/bav-demo/proto/society/go_pb"

	pbAuthen "github.com/bav-demo/proto/authenticate/go_pb"

	SocietyController "github.com/bav-demo/internal/controllers/society"
	"github.com/bav-demo/internal/middlewares"
)

type Route struct {
	pb.SocietyServiceServer
	pbAuthen.AuthenticationServiceServer
}

func (r *Route) NewRoute() *grpc.Server {

	// init server with Unary gRPC middleware
	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			middlewares.AuthInterceptor,
			middlewares.UnaryServerInterceptor,
		),
	)

	pb.RegisterSocietyServiceServer(grpcServer, r)
	pbAuthen.RegisterAuthenticationServiceServer(grpcServer, r)

	return grpcServer
}

func (s *Route) GetAllPostRPC(ctx context.Context, in *pb.RequestPosts) (*pb.ResponsePosts, error) {
	return SocietyController.GetAllPostRPC(ctx, in)
}

func (s *Route) GetPostById(ctx context.Context, in *pb.RequestPost) (*pb.ResponsePost, error) {

	return SocietyController.GetPostById(ctx, in)
}

func (s *Route) GetAllCommentsFromPost(ctx context.Context, in *pb.RequestPost) (*pb.ResponseComments, error) {
	return SocietyController.GetAllCommentsFromPost(ctx, in)
}

func (s *Route) AddPost(ctx context.Context, in *pb.RequestAddPost) (*pb.ResponseUpdate, error) {
	return SocietyController.AddPost(ctx, in)
}

func (s *Route) AddComment(ctx context.Context, in *pb.RequestAddComment) (*pb.ResponseUpdate, error) {
	return SocietyController.AddComment(ctx, in)
}

func (s *Route) UpdateComment(ctx context.Context, in *pb.RequestUpdateComment) (*pb.ResponseUpdate, error) {
	return SocietyController.UpdateComment(ctx, in)
}
