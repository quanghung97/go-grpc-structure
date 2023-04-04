package repositories

import (
	entity "github.com/bav-demo/internal/models/society"
	pb "github.com/bav-demo/proto/society/go_pb"
)

type ISocietyRespository interface {
	FindAllPost() ([]*pb.Post, error)
	FindPostById(postId string) (*pb.Post, error)
	FindAllCommentFromPosts(postId string) ([]*pb.Comment, error)
	AddComment(post entity.Comment) error
	AddPost(post entity.Post) error
	UpdateComment(commentId string, content string) error
}
