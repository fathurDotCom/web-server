package routes

import (
	"net/http"
	"web-server/controllers"
	"web-server/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	userGroup := router.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.POST("/", middleware.AuthRequired, controllers.CreateUsers)
	}

	return router
}
