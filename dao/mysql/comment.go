package mysql

import "go_bbs/models"

// CreateComment 创建评论
func CreateComment(p *models.ParamComment, pictureURL string) (err error) {
	sqlStr := `insert into comment(
    author_id, content, post_id, create_time, username, picture_url)
    values (?, ?, ?, ?, ?)
    `
	_, err = db.Exec(sqlStr, p.AuthorID, p.Content, p.PostId, p.CreateTime, p.UserName, pictureURL)
	return
}

// GetCommentFromPostId 根据帖子id查询出帖子评论
func GetCommentFromPostId(postId int64) (comment []*models.Comment, err error) {
	sqlStr := `select
	username, content, create_time, picture_url
	from comment
	where post_id = ?
	order by create_time
	desc
    `
	err = db.Select(&comment, sqlStr, postId)
	return
}
