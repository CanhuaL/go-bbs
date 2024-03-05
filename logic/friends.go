package logic

import (
	"go.uber.org/zap"
	"go_bbs/dao/mysql"
	"go_bbs/models"
	"time"
)

func AddFriend(p *models.ParamFriendRelation) (err error) {
	// 判断是否已经是好友，不是好友则继续往下
	if err = mysql.CheckFriendExist(p.UserId, p.FriendId); err != nil {
		return err
	}
	friend := &models.Friend{
		UserId:    p.UserId,
		FriendId:  p.FriendId,
		Status:    p.Status,
		CreatedAt: time.Now(),
	}
	return mysql.AddFriend(friend)
}

func ListFriend(p *models.ParamFriendRelation) (friendList []*models.FriendList, err error) {
	friendList, err = mysql.ListFriend(p.UserId)
	return
}

func DeleteFriend(p *models.ParamFriendRelation) (err error) {
	// 判断是否已经是好友，是则执行删除操作
	if err = mysql.CheckFriendExist(p.UserId, p.FriendId); err != nil {
		return mysql.DeleteFriend(p.UserId, p.FriendId)
	}
	zap.L().Error("删除的用户不存在！")
	return
}

func ConfirmFriend(p *models.ParamConfirmFriend) (err error) {
	return mysql.ConfirmFriend(p.UserId, p.FriendId, p.Status)
}
