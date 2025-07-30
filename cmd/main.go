package main

import (
	"fmt"

	"github.com/itisroach/go-blog/config"
	"github.com/itisroach/go-blog/database"
	"github.com/itisroach/go-blog/migration"
)


func init() {
	config.LoadEnvVariables()
	database.ConnectDatabase()
	migration.MakeMigrations()
}

func main() {
	fmt.Println("Hello World")
}