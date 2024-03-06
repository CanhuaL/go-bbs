package models

type PrivateChat struct {
	SenderId   int64  `db:"sender_id"`
	ReceiverId int64  `db:"receiver_id"`
	Content    string `db:"content"`
}
