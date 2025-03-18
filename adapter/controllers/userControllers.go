package controllers

import (
	repository "gin_api/adapter/Repository"
	"gin_api/adapter/handlers"
	service "gin_api/adapter/servise"
	"gin_api/core/commands"
	"gin_api/core/events"
	"gin_api/core/interfaces"
	"gin_api/initializer"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userCommandHandler *handlers.UserCommandHandler
var userReadService interfaces.UserReadService

// InitUserController initializes the user controller with Redis publisher
func InitUserController() {
	// Set up Redis publisher for events
	publisher := events.NewRedisEventPublisher(initializer.RedisClient, "user-events")
	userCommandHandler = handlers.NewUserCommandHandler(publisher)

	// Set up read service
	userReadRepo := repository.NewMongoUserRepository()
	userReadService = service.NewUserReadService(userReadRepo)
}

func CreateUser(c *gin.Context) {
	// Check if handler is initialized
	if userCommandHandler == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User command handler not initialized",
		})
		return
	}

	var body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Create command
	command := commands.NewCreateUserCommand(body.Name, body.Email)

	// Handle command
	err := userCommandHandler.Handle(command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "User creation command processed",
		"command_id": command.GetID(),
	})
}

func GetAll(c *gin.Context) {
	// Check if service is initialized
	if userReadService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User read service not initialized",
		})
		return
	}

	users, err := userReadService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetByID(c *gin.Context) {
	// Check if service is initialized
	if userReadService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User read service not initialized",
		})
		return
	}

	userID := c.Param("id")
	user, err := userReadService.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
