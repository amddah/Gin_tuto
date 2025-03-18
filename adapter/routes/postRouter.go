package routes

import (
	"github.com/gin-gonic/gin"
	"gin_api/adapter/controllers"
)

func PostRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/posts")
	{
		v1.POST("/", controllers.PostCreate)
		v1.GET("/", controllers.PostIndex)
		v1.GET("/:id", controllers.PostShow)
		v1.PUT("/:id", controllers.PostUpdate)
		v1.DELETE("/:id", controllers.PostDelete)
	}
}
