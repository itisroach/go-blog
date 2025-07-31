package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itisroach/go-blog/controllers"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/auth/register", controllers.RegisterUser)
		api.POST("/auth/login", controllers.LoginUser)
	}

	return router

}