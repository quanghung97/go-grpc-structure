package society

import (
	"context"
	"log"

	entity "github.com/bav-demo/micro1/internal/models/society"
	initConnect "github.com/bav-demo/micro1/server/init"

	util "github.com/bav-demo/micro1/pkg/uuid"

	repositories "github.com/bav-demo/micro1/internal/repositories/domain1/interface"
	pb "github.com/bav-demo/micro1/proto/society/go_pb"
)

var connection = initConnect.GetInstance()

var societyResponse = repositories.NewSocietyRespository()

func GetAllPostRPC(ctx context.Context, in *pb.RequestPosts) (*pb.ResponsePosts, error) {
	log.Printf("Receive message body from client: %v", in)

	result, err := societyResponse.FindAllPost()
	if err != nil {
		return nil, err
	}

	return &pb.ResponsePosts{Posts: result}, nil
}

func GetPostById(ctx context.Context, in *pb.RequestPost) (*pb.ResponsePost, error) {

	log.Println("Receive message body from client: %v", in)

	result, err := societyResponse.FindPostById(in.PostId)
	if err != nil {
		return nil, err
	}

	return &pb.ResponsePost{Post: result}, nil
}

func GetAllCommentsFromPost(ctx context.Context, in *pb.RequestPost) (*pb.ResponseComments, error) {

	log.Printf("Receive message body from client: %v", in)

	result, err := societyResponse.FindAllCommentFromPosts(in.PostId)
	if err != nil {
		return nil, err
	}
	return &pb.ResponseComments{Comments: result}, nil
}

func AddPost(ctx context.Context, in *pb.RequestAddPost) (*pb.ResponseUpdate, error) {

	log.Printf("Receive message body from client: %v", in)

	var post entity.Post = entity.Post{
		PostId: util.GenUuid(),
		Title:  in.Title,
	}

	err := societyResponse.AddPost(post)
	if err != nil {
		return nil, err
	}
	return &pb.ResponseUpdate{Status: 200}, nil
}

func AddComment(ctx context.Context, in *pb.RequestAddComment) (*pb.ResponseUpdate, error) {

	log.Printf("Receive message body from client: %v", in)

	var post entity.Comment = entity.Comment{
		CommentId: util.GenUuid(),
		PostId:    in.PostId,
		Content:   in.Content,
	}

	err := societyResponse.AddComment(post)
	if err != nil {
		return nil, err
	}

	return &pb.ResponseUpdate{Status: 200}, nil
}

func UpdateComment(ctx context.Context, in *pb.RequestUpdateComment) (*pb.ResponseUpdate, error) {

	log.Printf("Receive message body from client: %v", in)

	err := societyResponse.UpdateComment(in.CommentId, in.Content)
	if err != nil {
		return nil, err
	}

	return &pb.ResponseUpdate{Status: 200}, nil
}
