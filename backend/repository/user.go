package repository

import (
	"todo-go/models"

	"github.com/google/uuid"
)

type UserRepo interface {
    CreateUser(*models.User) error
    ReadAllUsers() ([]models.User, error)
    ReadUserByID(uuid.UUID) (*models.User, error)
    UpdateUser(*models.User) error
    DeleteUser(uuid.UUID) error
}
