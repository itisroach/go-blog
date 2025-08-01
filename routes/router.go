package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itisroach/go-blog/controllers"
	"github.com/itisroach/go-blog/middlewares"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/auth/register", controllers.RegisterUser)
 		api.POST("/auth/login", controllers.LoginUser)
		
		protected := api.Group("/")
		protected.Use(middlewares.AuthMiddleware())
		
		protected.POST("/auth/refresh", controllers.RefreshToken)

		protected.GET("/users/:username", controllers.GetUser)

		protected.POST("/posts/write", controllers.NewPost)
	}

	return router

}