package services

import (
	"net/http"

	"github.com/itisroach/go-blog/models"
	"github.com/itisroach/go-blog/repositories"
	"github.com/itisroach/go-blog/utils"
)

func LoginService(reqBody *models.LoginRequest) (*models.JWTResponse, *utils.CustomError) {

	user, err, password := repositories.GetUser(reqBody.Username, true)

	if err != nil {
		return nil, &utils.CustomError{
			Code:    http.StatusBadRequest,
			Message: "username or password is wrong",
		}
	}

	isCorrect, _ := utils.ComparePassword(reqBody.Password, password)

	if !isCorrect {
		return nil, &utils.CustomError{
			Code:    http.StatusBadRequest,
			Message: "username or password is wrong",
		}
	}

	access, _ := utils.GenerateJWTToken(user.Username, false)
	refresh, err := utils.GenerateJWTToken(user.Username, true)

	if err != nil {
		return nil, &utils.CustomError{
			Code:    http.StatusInternalServerError,
			Message: "it's not your fault, something went wrong",
		}
	}

	return &models.JWTResponse{
		Access:  access,
		Refresh: refresh,
	}, nil
}