package database

import (
	"os"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func InitDB(log *zap.Logger) *pgx.Conn {
	log.Info("Connecting to database")
	log.Info(os.Getenv("DATABASE_URL"))
	database, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error connecting to database", zap.Error(err))
	}
	return database
}
