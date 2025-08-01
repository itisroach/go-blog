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
	
	var err error

	if reqBody.Password, err = utils.HashString(reqBody.Password); err != nil {

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



func GetUserService(username string) (*models.UserResponse, *utils.CustomError) {

	userObj, err , _ := repositories.GetUser(username, false)


	if err != nil {
		return nil, &utils.CustomError{
			Code: 404,
			Message: err.Error(),
		}
	}

	return userObj, nil
}

