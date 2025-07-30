package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/itisroach/go-blog/models"
	"github.com/itisroach/go-blog/repositories"
	"github.com/itisroach/go-blog/utils"
)

func RegisterUser(c *gin.Context) {

	var reqBody *models.User

	// check for errors in user body
	if err := c.ShouldBindJSON(&reqBody); err != nil {

		// using field error validator from gorm
		var validator validator.ValidationErrors

		// convert the error to a user friendly error
		if errors.As(err, &validator) {
			out := make([]string, len(validator))

			for i, fieldError := range validator {
				out[i] = utils.FormatError(fieldError)
			}
			
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": out,
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}


	userFound, _ := repositories.GetUser(reqBody.Username)

	if userFound != nil {

		c.JSON(http.StatusConflict, gin.H{
			"error": "username is already taken",
		})

		return
	}

	if err := reqBody.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "it's not your fault, something went wrong",
		})

		return
	}


	if err := repositories.CreateUser(reqBody); err != nil {
		log.Fatal(err)
		
	}


	
	c.JSON(http.StatusCreated, reqBody.NewUserResponse())

}