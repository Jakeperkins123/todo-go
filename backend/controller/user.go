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

type UserController struct {
	log  *zap.Logger
	repo repository.UserRepo
}

func NewUserController(log *zap.Logger, db *pgx.Conn) *UserController {
	userRepo := repository.NewUserRepoImpl(log, db)
	return &UserController{
		log:  log,
		repo: userRepo,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		uc.log.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON"})
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		uc.log.Error("Error generating UUID", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating UUID"})
		return
	}
	user.ID = id

	err = uc.repo.CreateUser(&user)
	if err != nil {
		uc.log.Error("Error creating user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}
	uc.log.Info("User created successfully", zap.String("id", user.ID.String()))
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "id": user.ID})
}

func (uc *UserController) ReadAllUsers(c *gin.Context) {
	users, err := uc.repo.ReadAllUsers()
	if err != nil {
		uc.log.Error("Error reading users", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading users"})
		return
	}
	uc.log.Info("Users read successfully")
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *UserController) ReadUserByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		uc.log.Error("Error parsing UUID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing UUID"})
		return
	}
	user, err := uc.repo.ReadUserByID(id)
	if err != nil {
		uc.log.Error("Error reading user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading user"})
		return
	}
	uc.log.Info("User read successfully", zap.String("id", user.ID.String()))
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		uc.log.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON"})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		uc.log.Error("Error parsing UUID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing UUID"})
		return
	}
	user.ID = id

	err = uc.repo.UpdateUser(&user)
	if err != nil {
		uc.log.Error("Error updating user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}
	uc.log.Info("User updated successfully", zap.String("id", user.ID.String()))
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "id": user.ID})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		uc.log.Error("Error parsing UUID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing UUID"})
		return
	}

	err = uc.repo.DeleteUser(id)
	if err != nil {
		uc.log.Error("Error deleting user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}
	uc.log.Info("User deleted successfully", zap.String("id", id.String()))
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "id": id})
}
