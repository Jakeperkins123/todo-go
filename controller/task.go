package controller

import (
	"net/http"
	"todo-go/models"
	"todo-go/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type TaskController struct {
	log  *zap.Logger
	repo repository.TaskRepoImpl
}

func NewTaskController(logger *zap.Logger, db *pgx.Conn) *TaskController {
	taskRepo := repository.NewTaskRepoImpl(db)
	return &TaskController{
		log:  logger,
		repo: taskRepo,
	}
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		tc.log.Error("Error while binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		tc.log.Error("Error while generating UUID", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	task.ID = id

	err = tc.repo.CreateTask(&task)
	if err != nil {
		tc.log.Error("Error while creating task", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.log.Info("Task created successfully", zap.String("id", task.ID.String()))
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully", "id": task.ID})
}

func (tc *TaskController) ReadAllTasks(c *gin.Context) {
	tasks, err := tc.repo.ReadAllTasks()
	if err != nil {
		tc.log.Error("Error while reading tasks", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.log.Info("Tasks read successfully")
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (tc *TaskController) ReadTaskByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		tc.log.Error("Error while parsing UUID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.repo.ReadTaskByID(id)
	if err != nil {
		tc.log.Error("Error while reading task", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.log.Info("Task read successfully", zap.String("id", task.ID.String()))
	c.JSON(http.StatusOK, gin.H{"task": task})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	var task models.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		tc.log.Error("Error while binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		tc.log.Error("Error while parsing UUID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.ID = id

	err = tc.repo.UpdateTask(&task)
	if err != nil {
		tc.log.Error("Error while updating task", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.log.Info("Task updated successfully", zap.String("id", task.ID.String()))
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully", "id": task.ID})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		tc.log.Error("Error while parsing UUID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = tc.repo.DeleteTask(id)
	if err != nil {
		tc.log.Error("Error while deleting task", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tc.log.Info("Task deleted successfully", zap.String("id", id.String()))
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully", "id": id})
}
