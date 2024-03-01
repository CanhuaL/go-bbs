package models

import (
	"time"
)

// 定义请求的参数结构体

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	RePassword string `json:"re_password"`
	Gender     int32  `json:"gender"`
	Avatar     []byte `json:"avatar"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

// ParamVoteData 投票数据
type ParamVoteData struct {
	// UserID 从请求中获取当前的用户
	PostID    string `json:"post_id" binding:"required"`               // 贴子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1" ` // 赞成票(1)还是反对票(-1)取消投票(0)
}

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page" example:"1"`       // 页码
	Size        int64  `json:"size" form:"size" example:"10"`      // 每页数据量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}

// ParamComment 创建评论参数
type ParamComment struct {
	PostId     int64     `json:"post_id" db:"post_id"` //帖子Id
	UserName   string    `json:"user_name" db:"username" binding:"required"`
	AuthorID   int64     `json:"author_id" db:"author_id" binding:"required"` //评论者id
	Content    string    `json:"content" db:"content" binding:"required"`     //内容
	CreateTime time.Time `json:"create_time" db:"create_time"`                //创建时间
	PictureURL string    `json:"picture_url" db:"picture_url"`
}
