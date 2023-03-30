package routing

import (
	"context"
	// "github.com/bav-demo/micro1/internal/controllers/society"

	"google.golang.org/grpc"

	pb "github.com/bav-demo/micro1/proto/society/go_pb"

	pbAuthen "github.com/bav-demo/micro1/proto/authenticate/go_pb"

	SocietyController "github.com/bav-demo/micro1/internal/controllers/society"
	"github.com/bav-demo/micro1/internal/middlewares"
)

type Route struct {
	pb.SocietyServiceServer
	pbAuthen.AuthenticationServiceServer
}

func (r *Route) NewRoute() *grpc.Server {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middlewares.AuthInterceptor))

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
