package domain

import "time"

type Activity struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Title     string    `json:"title" binding:"required" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
