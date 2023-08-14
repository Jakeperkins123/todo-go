package repository

import (
	"context"
	"time"
	"todo-go/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type UserRepoImpl struct {
	log *zap.Logger
	db  *pgx.Conn
}

func NewUserRepoImpl(log *zap.Logger, db *pgx.Conn) *UserRepoImpl {
	return &UserRepoImpl{
		log: log,
		db:  db,
	}
}

func (u *UserRepoImpl) CreateUser(user *models.User) error {
	ctx := context.Background()
	_, err := u.db.Exec(ctx, "INSERT INTO USERS (id, username, email) VALUES ($1, $2, $3)", user.ID, user.Username, user.Email)
	return err
}

func (u *UserRepoImpl) ReadAllUsers() ([]models.User, error) {
	var users []models.User
	ctx := context.Background()
	rows, err := u.db.Query(ctx, "select id, username, email, created_at, updated_at, deleted_at FROM USERS where deleted_at is null")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepoImpl) ReadUserByID(id uuid.UUID) (*models.User, error) {
    ctx := context.Background()
    stmt := `SELECT id, username, email, created_at, updated_at, deleted_at FROM USERS WHERE id = $1 AND deleted_at IS NULL`
    var user models.User
    err := u.db.QueryRow(ctx, stmt, id).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
    return &user, err
}

func (u *UserRepoImpl) UpdateUser(user *models.User) error {
    ctx := context.Background()
    _, err := u.db.Exec(ctx, "UPDATE USERS SET username = $1, email = $2, updated_at = $3 WHERE id = $4", &user.Username, &user.Email, time.Now(), &user.ID)
    return err
}

func (u *UserRepoImpl) DeleteUser(id uuid.UUID) error {
    ctx := context.Background()
    _, err := u.db.Exec(ctx, "UPDATE USERS SET deleted_at = $1 WHERE id = $2", time.Now(), id)
    return err
}
