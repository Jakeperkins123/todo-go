package main

import (
	"todo-go/database"
	"todo-go/logger"
	"todo-go/router"
)

func main() {
	log := logger.InitLogger()
	defer log.Sync()
	db := database.InitDB(log)
	router.InitRouter(log, db)
}
