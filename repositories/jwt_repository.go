package repositories

import (
	"errors"
	"time"

	"github.com/itisroach/go-blog/database"
	"github.com/itisroach/go-blog/models"
	"gorm.io/gorm"
)

func SaveTokenToDB(jwtToken string, username string, expTime time.Time) error {
	
	token := models.RefreshToken{
		Token: jwtToken,
		Username: username,
		ExpiresAt: expTime,
	}
	
	err := DeleteToken(username)
	
	if err != nil {
		return err
	}

	result := database.DB.Create(&token)

	if result.Error != nil {
		return result.Error
	}
	
	return nil
}

func DeleteToken(username string) error {
	result := database.DB.Delete(&models.RefreshToken{}, "username = ?", username)

	if result.Error != nil {
		return result.Error
	}

	return nil
}


func IsTokenValidInDB(token string) (bool, error) {

	var tokenInstance *models.RefreshToken

	result := database.DB.First(&tokenInstance, "token = ?", token)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, errors.New("token is not valid")
	}

	if result.Error != nil {
		return false, result.Error
	}


	expired := time.Now().After(tokenInstance.ExpiresAt)

	if expired {
		return false, errors.New("refresh token is expired")
	}

	return true, nil
}