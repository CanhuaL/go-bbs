package models

import "time"

type Comment struct {
	CommentId  int64     `json:"comment_id" db:"comment_id"`   //评论的唯一标识
	AuthorID   int64     `json:"author_id" db:"author_id"`     //评论者id
	AuthorName int64     `json:"author_name" db:"author_name"` //评论者id
	UserName   string    `json:"user_name" db:"username"`      //评论者昵称
	Content    string    `json:"content" db:"content"`         //内容
	PictureURL string    `json:"picture_url" db:"picture_url"` //评论照片的url
	CreateTime time.Time `json:"create_time" db:"create_time"` //创建时间
}
