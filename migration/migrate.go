package migration

import (
	"log"
	"os"

	"github.com/itisroach/go-blog/models"
	"github.com/itisroach/go-blog/database"
)

func MakeMigrations() {

	if err := database.DB.AutoMigrate(&models.User{}, &models.RefreshToken{}); err != nil {
		log.Fatal("migrations failed")
		os.Exit(1)
	}

}