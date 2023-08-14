package repository

import (
	"context"
	"time"
	"todo-go/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TaskRepoImpl struct {
	db *pgx.Conn
}

func NewTaskRepoImpl(db *pgx.Conn) TaskRepoImpl {
	return TaskRepoImpl{db: db}
}

func (t *TaskRepoImpl) CreateTask(task *models.Task) error {
	stmt := `INSERT INTO TASKS (id, title, description, user_id) VALUES ($1, $2, $3, $4)`
	_, err := t.db.Exec(context.Background(), stmt, task.ID, task.Title, task.Description, task.UserID)
	return err
}

func (t *TaskRepoImpl) ReadAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	stmt := `SELECT id, title, description, user_id, created_at, completed_at, updated_at, deleted_at FROM TASKS WHERE deleted_at IS NULL`
	rows, err := t.db.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.UserID, &task.CreatedAt, &task.CompletedAt, &task.UpdatedAt, &task.DeletedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

    if err := rows.Err(); err != nil {
        return nil, err
    }

	return tasks, nil
}

func (t *TaskRepoImpl) ReadTaskByID(id uuid.UUID) (*models.Task, error){
    var task models.Task
    stmt := `SELECT id, title, description, user_id, created_at, completed_at, updated_at, deleted_at FROM TASKS WHERE id = $1 AND deleted_at IS NULL`
    err := t.db.QueryRow(context.Background(), stmt, id).Scan(&task.ID, &task.Title, &task.Description, &task.UserID, &task.CreatedAt, &task.CompletedAt, &task.UpdatedAt, &task.DeletedAt)
    if err != nil {
        return nil, err
    }
    return &task, nil
}

func (t *TaskRepoImpl) UpdateTask(task *models.Task) error {
    stmt := `UPDATE TASKS SET title = $1, description = $2, completed_at = $3, updated_at = $4 WHERE id = $5 AND deleted_at IS NULL`
    _, err := t.db.Exec(context.Background(), stmt, task.Title, task.Description, task.CompletedAt, time.Now(), task.ID)
    if err != nil {
        return err
    }
    return nil
}

func (t *TaskRepoImpl) DeleteTask(id uuid.UUID) error {
    stmt := `UPDATE TASKS SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`
    _, err := t.db.Exec(context.Background(), stmt, time.Now(), id)
    if err != nil {
        return err
    }
    return nil
}

