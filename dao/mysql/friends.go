package mysql

import (
	"go_bbs/models"
)

func AddFriend(friend *models.Friend) (err error) {
	sqlStr := `insert into friend_relationship
	(user_id, friend_id, status, created_at) values (?,?,?,?)
    `
	_, err = db.Exec(sqlStr, friend.UserId, friend.FriendId, friend.Status, friend.CreatedAt)
	return
}

func CheckFriendExist(userId, friendId int64) (err error) {
	sqlStr := `select count(*)
	from friend_relationship
	where user_id = ? and friend_id = ?
    `
	var count int
	if err = db.Get(&count, sqlStr, userId, friendId); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func ListFriend(userId int64) (friendList []*models.FriendList, err error) {
	sqlStr := `select friend_id
	from friend_relationship
	where user_id = ?
	order by created_at 
	desc
    `
	err = db.Select(&friendList, sqlStr, userId)
	return
}

func DeleteFriend(userId, friendId int64) (err error) {
	sqlStr := `delete  
	from friend_relationship
	where user_id = ? and friend_id = ?
	`
	_, err = db.Exec(sqlStr, userId, friendId)
	if err != nil {
		return err
	}
	return
}

func ConfirmFriend(userId, friendId int64, status string) (err error) {
	sqlStr := `update friend_relationship
	set status = ?
	where (user_id = ? and friend_id = ?)
	`
	_, err = db.Exec(sqlStr, status, userId, friendId)
	if err != nil {
		return err
	}
	return
}
