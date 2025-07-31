package repositories

import (
	"github.com/itisroach/go-blog/database"
	"github.com/itisroach/go-blog/models"
)

func CreateUser(user *models.User) error {
	
	result := database.DB.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}


func GetUser(username string) (*models.UserResponse, error) {
	
	var user *models.User

	result := database.DB.First(&user, "username = ?", username) 

	if result.Error != nil {
		return nil, result.Error
	}

	return models.NewUserResponse(user), nil
}