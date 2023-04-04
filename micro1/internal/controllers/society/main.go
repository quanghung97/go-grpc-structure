package society

import (
	"context"
	"fmt"

	entity "github.com/bav-demo/internal/models/society"
	initConnect "github.com/bav-demo/server/init"

	util "github.com/bav-demo/pkg/uuid"

	repositories "github.com/bav-demo/internal/repositories/domain1/interface"
	pb "github.com/bav-demo/proto/society/go_pb"
)

var connection = initConnect.GetInstance()

var societyResponse = repositories.NewSocietyRespository()

/**
 * GetAllPostRPC
 * ctx
 * in
 */
func GetAllPostRPC(ctx context.Context, in *pb.RequestPosts) (*pb.ResponsePosts, error) {
	fmt.Printf("Receive message body from client: %v", in)

	result, err := societyResponse.FindAllPost()
	if err != nil {
		return nil, err
	}
	return &pb.ResponsePosts{Posts: result}, nil
}

func GetPostById(ctx context.Context, in *pb.RequestPost) (*pb.ResponsePost, error) {

	fmt.Println("Receive message body from client: %v", in)

	result, err := societyResponse.FindPostById(in.PostId)
	if err != nil {
		return nil, err
	}

	return &pb.ResponsePost{Post: result}, nil
}

func GetAllCommentsFromPost(ctx context.Context, in *pb.RequestPost) (*pb.ResponseComments, error) {

	fmt.Printf("Receive message body from client: %v", in)

	result, err := societyResponse.FindAllCommentFromPosts(in.PostId)
	if err != nil {
		return nil, err
	}
	return &pb.ResponseComments{Comments: result}, nil
}

func AddPost(ctx context.Context, in *pb.RequestAddPost) (*pb.ResponseUpdate, error) {

	fmt.Printf("Receive message body from client: %v", in)

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

	fmt.Printf("Receive message body from client: %v", in)

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

	fmt.Printf("Receive message body from client: %v", in)

	err := societyResponse.UpdateComment(in.CommentId, in.Content)
	if err != nil {
		return nil, err
	}

	return &pb.ResponseUpdate{Status: 200}, nil
}
