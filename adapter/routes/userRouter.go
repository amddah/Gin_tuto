package routes

import (
	"gin_api/adapter/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	v2 := r.Group("/users")
	{
		v2.POST("/", controllers.CreateUser)
		v2.GET("/", controllers.GetAll)
		v2.GET("/:id", controllers.GetByID)
	}
}
