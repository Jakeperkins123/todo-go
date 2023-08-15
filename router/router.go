package router

import (
	"todo-go/controller"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func InitRouter(log *zap.Logger, db *pgx.Conn) {
	taskController := controller.NewTaskController(log, db)
	userController := controller.NewUserController(log, db)

	r := gin.Default()

	r.POST("/tasks", taskController.CreateTask)
	r.GET("/tasks", taskController.ReadAllTasks)
	r.GET("/tasks/:id", taskController.ReadTaskByID)
	r.PUT("/tasks/:id", taskController.UpdateTask)
	r.DELETE("/tasks/:id", taskController.DeleteTask)

	r.POST("/users", userController.CreateUser)
	r.GET("/users", userController.ReadAllUsers)
	r.GET("/users/:id", userController.ReadUserByID)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Error while starting server", zap.Error(err))
	}
}
