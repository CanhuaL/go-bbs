package mysql

import "go_bbs/models"

func SavePrivateChat(m *models.PrivateChat) (err error) {
	sqlStr := `insert into chat_private_messages
	(sender_id, receiver_id, content) 
	VALUES (?, ?, ?)
    `
	_, err = db.Exec(sqlStr, m.SenderId, m.ReceiverId, m.Content)
	return
}

func CheckPrivateChat(senderId, receiverId, page, size int64) (msg []*models.PrivateChat, err error) {
	sqlStr := `select
	sender_id, receiver_id, content
	from chat_private_messages
	where sender_id = ? and receiver_id
	order by sent_at
	desc 
	limit ?,?
	`
	msg = make([]*models.PrivateChat, 0, 2)
	err = db.Select(&msg, sqlStr, senderId, receiverId, (page-1)*size, size)
	return
}
