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


func GetUser(username string, needPassword bool) (*models.UserResponse, string, error) {
	
	var user *models.User

	result := database.DB.First(&user, "username = ?", username) 

	if result.Error != nil {
		return nil, "", result.Error
	}

	if needPassword {
		return models.NewUserResponse(user), user.Password, nil
	}

	return models.NewUserResponse(user), "", nil
}

func GetUserRawData(username string) (*models.User, error) {
	var user *models.User

	result := database.DB.First(&user, "username = ?", username) 

	if result.Error != nil {
		return nil, result.Error
	}


	return user, nil
}