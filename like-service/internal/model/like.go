package model

import "time"

type Like struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id"`
	MessageID int       `json:"message_id"`
	CreatedAt time.Time `json:"created_at"`
}
