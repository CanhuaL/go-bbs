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
