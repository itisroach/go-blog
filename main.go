// @title Blog API
// @version 1.0
// @description This is a sample server for a blog API.
// @BasePath /api
package main

import (
	"os"

	"github.com/itisroach/go-blog/config"
	"github.com/itisroach/go-blog/database"
	"github.com/itisroach/go-blog/migration"
	"github.com/itisroach/go-blog/routes"

	_ "github.com/itisroach/go-blog/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func init() {
	config.LoadEnvVariables()
	database.ConnectDatabase()
	migration.MakeMigrations()
}

func main() {
	
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}


	router := routes.SetupRouter()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":" + port)

}