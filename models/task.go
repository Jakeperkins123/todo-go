package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	UserID      uuid.UUID  `json:"user_id"`
	CompletedAt *time.Time `json:"completed_at" `
	DeletedAt   *time.Time `json:"deleted_at" `
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
