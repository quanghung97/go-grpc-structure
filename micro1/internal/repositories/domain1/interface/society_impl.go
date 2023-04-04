package repositories

import (
	"errors"

	entity "github.com/bav-demo/internal/models/society"
	pb "github.com/bav-demo/proto/society/go_pb"
	initConnect "github.com/bav-demo/server/init"
)

var connection = initConnect.GetInstance()

type SocietyRespository struct{}

func NewSocietyRespository() ISocietyRespository {
	return &SocietyRespository{}
}

func (s *SocietyRespository) FindAllPost() ([]*pb.Post, error) {
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

	return posts_, nil
}
func (s *SocietyRespository) FindPostById(postId string) (*pb.Post, error) {
	var postQuery *entity.Post

	if connection.Db.Where("post_id = ?", postId).Find(&postQuery).Error != nil {
		return nil, errors.New("post not found")
	}

	post := pb.Post{
		PostId: postQuery.PostId,
		Title:  postQuery.Title,
	}

	return &post, nil
}
func (s *SocietyRespository) FindAllCommentFromPosts(postId string) ([]*pb.Comment, error) {
	comments := make([]*entity.Comment, 0)
	comments_ := make([]*pb.Comment, 0)

	if connection.Db.Where("post_id = ?", postId).Find(&comments).Error != nil {
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
	return comments_, nil
}
func (s *SocietyRespository) AddComment(post entity.Comment) error {

	if connection.Db.Create(&post).Error != nil {
		return errors.New("something went wrong")
	}
	return nil
}
func (s *SocietyRespository) AddPost(post entity.Post) error {
	if connection.Db.Create(&post).Error != nil {
		return errors.New("something went wrong")
	}
	return nil
}

func (s *SocietyRespository) UpdateComment(commentId string, content string) error {
	var post entity.Comment
	if connection.Db.Where("comment_id = ?", commentId).Find(&post).Error != nil {
		return errors.New("comment not found")
	}

	post.Content = content

	if connection.Db.Where("comment_id = ?", commentId).Save(&post).Error != nil {
		return errors.New("something went wrong")
	}
	return nil
}
