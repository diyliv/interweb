package models

import "time"

type User struct {
	UserId            int       `json:"user_id"`
	UserTelegramId    int64     `json:"user_telegram_id"`
	UserFirstRequest  time.Time `json:"user_first_request"`
	UserRequestsCount int       `json:"user_requests_count"`
}
