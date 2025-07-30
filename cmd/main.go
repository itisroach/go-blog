package main

import (
	"os"

	"github.com/itisroach/go-blog/config"
	"github.com/itisroach/go-blog/database"
	"github.com/itisroach/go-blog/migration"
	"github.com/itisroach/go-blog/routes"
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

	router.Run(":" + port)

}