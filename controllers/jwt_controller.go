package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itisroach/go-blog/models"
	"github.com/itisroach/go-blog/services"
	"github.com/itisroach/go-blog/utils"

	_ "github.com/itisroach/go-blog/docs"
)

// LoginUser godoc
// @Summary      Login as an user
// @Description  Login and save jwt tokens
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.LoginRequest  true  "Login user data"
// @Success      201   {object}  models.JWTResponse
// @Failure      400   {object}  map[string]interface{}
// @Failure      409   {object}  map[string]string
// @Router       /auth/login [post]
func LoginUser(c *gin.Context) {

	var reqBody *models.LoginRequest
	
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


	result, err := services.LoginService(reqBody)


	if err != nil {
		c.JSON(err.Code, gin.H{
			"message": err.Message,
		})
		return
	}


	c.JSON(http.StatusOK, result)
	
}


// RefreshToken  godoc
// @Summary      Refresh access tokens
// @Description  Refresh access tokens that are expire so user can stay logged in
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.JWTRefreshRequest  true  "refresh token"
// @Success      201   {object}  models.JWTResponse
// @Failure      400   {object}  map[string]interface{}
// @Failure      409   {object}  map[string]string
// @Router       /auth/refresh [post]
func RefreshToken(c *gin.Context) {

	var reqBody *models.JWTRefreshRequest

	username, ok := c.Get("user")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "authorization header missed, make sure you have a valid token in headers",
		})
		return
	}


	if err := c.ShouldBindJSON(&reqBody); err != nil{
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


	response, err := services.RefreshTokenService(reqBody, username.(string))
	
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}


	c.JSON(http.StatusOK, response)

}