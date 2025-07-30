package main

import (
	"fmt"

	"github.com/itisroach/go-blog/config"
	"github.com/itisroach/go-blog/database"
)


func init() {
	config.LoadEnvVariables()
	database.ConnectDatabase()
}

func main() {
	fmt.Println("Hello World")
}