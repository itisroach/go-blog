package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itisroach/go-blog/models"
	"github.com/itisroach/go-blog/services"
	"github.com/itisroach/go-blog/utils"

	_ "github.com/itisroach/go-blog/docs"
)

// RegisterUser godoc
// @Summary      Register a new user
// @Description  Creates a new user account with a unique username
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.UserRequest  true  "User registration data"
// @Success      201   {object}  models.UserResponse
// @Failure      400   {object}  map[string]interface{}
// @Failure      409   {object}  map[string]string
// @Router       /auth/register [post]
func RegisterUser(c *gin.Context) {

	var reqBody *models.UserRequest

	// check for errors in user body
	if err := c.ShouldBindJSON(&reqBody); err != nil {

		allErrors := utils.GenerateUserFriendlyError(err)
		
		if allErrors == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": allErrors,
		})

		return
		
	}


	userInstance, err := services.CreateUser(reqBody)

	if err != nil {
		c.JSON(err.Code, gin.H{
			"message": err.Message,
		})
		return
	}
	
	c.JSON(http.StatusCreated, models.NewUserResponse(userInstance))

}



// GetUser godoc
// @Summary      Fetches a user
// @Description  Fetches a user by their username
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        username  path  	 string  true  "Username"
// @Success      200   	   {object}  models.UserResponse
// @Failure      404       {object}  map[string]interface{}
// @Router       /users/{username} [GET]
func GetUser(c *gin.Context) {

	username := c.Param("username")

	user, err := services.GetUserService(username)


	if err != nil {
		c.JSON(err.Code, gin.H{
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, user)
}