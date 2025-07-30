package main

import (
	"fmt"

	"github.com/itisroach/go-blog/config"
)


func init() {
	config.LoadEnvVariables()
}

func main() {
	fmt.Println("Hello World")
}