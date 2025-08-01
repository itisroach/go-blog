package repositories

import (
	"github.com/itisroach/go-blog/database"
	"github.com/itisroach/go-blog/models"
)

func CreatePost(post *models.Post) error {


	result := database.DB.Create(&post)


	if result.Error != nil {
		return result.Error
	}

	return nil

}