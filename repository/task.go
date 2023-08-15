package repository

import (
	"todo-go/models"

	"github.com/google/uuid"
)

type TaskRepo interface {
	CreateTask(*models.Task) error
	ReadAllTasks() ([]models.Task, error)
	ReadTaskByID(uuid.UUID) (*models.Task, error)
	UpdateTask(*models.Task) error
	DeleteTask(uuid.UUID) error
}
