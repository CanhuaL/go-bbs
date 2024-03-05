package models

import "time"

type Friend struct {
	UserId    int64     `db:"user_id"`
	FriendId  int64     `db:"friend_id"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"update_at"`
}

type FriendList struct {
	FriendId int64 `db:"friend_id"`
}
