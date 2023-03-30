package society

import (
	"context"
	"log"

	entity "github.com/bav-demo/micro1/internal/models/society"
	initConnect "github.com/bav-demo/micro1/server/init"

	"errors"

	util "github.com/bav-demo/micro1/pkg/uuid"

	pb "github.com/bav-demo/micro1/proto/society/go_pb"
)

var connection = initConnect.GetInstance()

func GetAllPostRPC(ctx context.Context, in *pb.RequestPosts) (*pb.ResponsePosts, error) {
	log.Printf("Receive message body from client: %v", in)

	posts := make([]*entity.Post, 0)
	posts_ := make([]*pb.Post, 0)

	if connection.Db.Find(&posts).Error != nil {
		return nil, errors.New("Societies not found")
	}

	for _, element := range posts {
		post := pb.Post{
			PostId: element.PostId,
			Title:  element.Title,
		}
		posts_ = append(posts_, &post)
	}

	return &pb.ResponsePosts{Posts: posts_}, nil
}

func GetPostById(ctx context.Context, in *pb.RequestPost) (*pb.ResponsePost, error) {

	log.Println("Receive message body from client: %v", in)

	var postQuery *entity.Post

	if connection.Db.Where("post_id = ?", in.PostId).Find(&postQuery).Error != nil {
		return nil, errors.New("post not found")
	}

	post := pb.Post{
		PostId: postQuery.PostId,
		Title:  postQuery.Title,
	}

	return &pb.ResponsePost{Post: &post}, nil
}

func GetAllCommentsFromPost(ctx context.Context, in *pb.RequestPost) (*pb.ResponseComments, error) {

	log.Printf("Receive message body from client: %v", in)

	comments := make([]*entity.Comment, 0)
	comments_ := make([]*pb.Comment, 0)

	if connection.Db.Where("post_id = ?", in.PostId).Find(&comments).Error != nil {
		return nil, errors.New("Societies not found")
	}

	for _, element := range comments {
		comment := pb.Comment{
			PostId:    element.PostId,
			Content:   element.Content,
			CommentId: element.CommentId,
		}
		comments_ = append(comments_, &comment)
	}

	return &pb.ResponseComments{Comments: comments_}, nil
}

func AddPost(ctx context.Context, in *pb.RequestAddPost) (*pb.ResponseUpdate, error) {

	log.Printf("Receive message body from client: %v", in)

	var post entity.Post = entity.Post{
		PostId: util.GenUuid(),
		Title:  in.Title,
	}

	if connection.Db.Create(&post).Error != nil {
		return nil, errors.New("something went wrong")
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

	if connection.Db.Create(&post).Error != nil {
		return nil, errors.New("something went wrong")
	}

	return &pb.ResponseUpdate{Status: 200}, nil
}

func UpdateComment(ctx context.Context, in *pb.RequestUpdateComment) (*pb.ResponseUpdate, error) {

	log.Printf("Receive message body from client: %v", in)

	var post entity.Comment

	if connection.Db.Where("comment_id = ?", in.CommentId).Find(&post).Error != nil {
		return nil, errors.New("comment not found")
	}

	post.Content = in.Content

	if connection.Db.Where("comment_id = ?", in.CommentId).Save(&post).Error != nil {
		return nil, errors.New("something went wrong")
	}

	return &pb.ResponseUpdate{Status: 200}, nil
}
