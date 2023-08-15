package database

import (
	"fmt"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func InitDB(log *zap.Logger) *pgx.Conn {
	log.Info("Connecting to database")
	// Construct the connection string
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=require",
		"postgres",
		"todo",
		"/cloudsql/todo-go-395906:us-central1:todo-go", 
		5432, // PostgreSQL port
		"todo-go",
	)
	database, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatal("Error connecting to database", zap.Error(err))
	}

	// Create the tables
	stmt := `-- Create the 'users' table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- Create the 'tasks' table
CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL ,
    user_id UUID NOT NULL REFERENCES users(id),
    completed_at TIMESTAMP,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
`
	_, err = database.Exec(context.Background(), stmt)
	if err != nil {
		log.Fatal("Error creating tables", zap.Error(err))
	}

	return database
}
