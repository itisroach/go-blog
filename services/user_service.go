package services

import (
	"log"
	"net/http"

	"github.com/itisroach/go-blog/models"
	"github.com/itisroach/go-blog/repositories"
	"github.com/itisroach/go-blog/utils"
)

func CreateUser(reqBody *models.UserRequest) (*models.User, *utils.CustomError) {
	userFound, _, _ := repositories.GetUser(reqBody.Username, false)

	if userFound != nil {

		return nil, &utils.CustomError{
			Code: http.StatusConflict,
			Message: "this username is already taken",
		}
	}

	if err := reqBody.HashPassword(); err != nil {

		return nil, &utils.CustomError{
			Code: http.StatusInternalServerError,
			Message: "it's not your fault, something went wrong",
		}
	}

	userInstance := reqBody.MakeUser()

	if err := repositories.CreateUser(userInstance); err != nil {
		log.Fatal(err)
	}

	return userInstance, nil
}



func LoginService(reqBody *models.LoginRequest) (*models.JWTResponse, *utils.CustomError) {

	user, err, password := repositories.GetUser(reqBody.Username, true)


	if err != nil {
		return nil, &utils.CustomError{
			Code: http.StatusBadRequest,
			Message: "username or password is wrong",
		}
	}

	
	isCorrect, _ := utils.ComparePassword(reqBody.Password, password)

	if !isCorrect {
		return nil, &utils.CustomError{
			Code: http.StatusBadRequest,
			Message: "username or password is wrong",
		}
	}


	access, _ := utils.GenerateJWTToken(user.Username, false)
	refresh, err := utils.GenerateJWTToken(user.Username, true)

	if err != nil {
		return nil, &utils.CustomError{
			Code: http.StatusInternalServerError,
			Message: "it's not your fault, something went wrong",
		}
	}

	return &models.JWTResponse{
		Access: access,
		Refresh: refresh,
	}, nil
}