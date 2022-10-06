package domain

import "time"

type Todo struct {
	ID              int       `json:"id"`
	ActivityGroupId string    `json:"activity_group_id" binding:"required"`
	Title           string    `json:"title" binding:"required"`
	IsActive        string    `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

type TodoUpdate struct {
	ID              int       `json:"id"`
	ActivityGroupId string    `json:"activity_group_id"`
	Title           string    `json:"title"`
	IsActive        string    `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}
