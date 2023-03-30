package society

type Post struct {
	PostId string `json:"post_id" db:"post_id" gorm:"primaryKey;`
	Title  string `json:"titel" db:"titel"`
}

type APIPostListAll struct {
	PostId string `json:"post_id" db:"post_id" gorm:"primaryKey;`
	Title  string `json:"titel" db:"titel"`
}

type Comment struct {
	CommentId string `json:"comment_id" gorm:"primaryKey;`
	Content   string `json:"content" db:"content"`
	PostId    string `json:"post_id" db:"post_id"`
}
