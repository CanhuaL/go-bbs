package models

import "time"

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Phone    string `db:"phone"`
	Email    string `db:"email"`
	Avatar   []byte `db:"avatar"`
	Gender   int32  `db:"gender"`
	Token    string
}

type SMS struct {
	SMSType    string    `db:"sms_type"`
	SMSContent string    `db:"sms_content"`
	Phone      string    `db:"phone"`
	SendTime   time.Time `db:"send_time"`
}
