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

var postCommandHandler *handlers.PostCommandHandler
var postReadService interfaces.PostReadService

// InitPostController initializes the post controller
func InitPostController() {
	// Set up Redis publisher for events
	publisher := events.NewRedisEventPublisher(initializer.RedisClient, "post-events")
	postCommandHandler = handlers.NewPostCommandHandler(publisher)

	// Setup read service
	postReadRepo := repository.NewMongoPostRepository()
	postReadService = service.NewPostReadService(postReadRepo)
}

func PostCreate(c *gin.Context) {
	// Check if handler is initialized
	if postCommandHandler == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Post command handler not initialized",
		})
		return
	}

	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Create command
	command := commands.NewCreatePostCommand(body.Title, body.Content)

	// Handle command
	err := postCommandHandler.Handle(command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Post creation command processed",
		"command_id": command.GetID(),
	})
}

func PostIndex(c *gin.Context) {
	// Check if service is initialized
	if postReadService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Post read service not initialized",
		})
		return
	}

	posts, err := postReadService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// Check if service is initialized
	if postReadService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Post read service not initialized",
		})
		return
	}

	postID := c.Param("id")
	post, err := postReadService.GetByID(postID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Check if handler is initialized
	if postCommandHandler == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Post command handler not initialized",
		})
		return
	}

	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	postID := c.Param("id")

	// Create command
	command := commands.NewUpdatePostCommand(postID, body.Title, body.Content)

	// Handle command
	err := postCommandHandler.Handle(command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Post update command processed",
		"command_id": command.GetID(),
	})
}

func PostDelete(c *gin.Context) {
	// Check if handler is initialized
	if postCommandHandler == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Post command handler not initialized",
		})
		return
	}

	postID := c.Param("id")

	// Create command
	command := commands.NewDeletePostCommand(postID)

	// Handle command
	err := postCommandHandler.Handle(command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Post deletion command processed",
		"command_id": command.GetID(),
	})
}
