package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itisroach/go-blog/models"
	"github.com/itisroach/go-blog/services"
	"github.com/itisroach/go-blog/utils"
)

// GetPosts godoc
// @Summary      Fetches all posts in database
// @Description  Fetches all posts in database but with limit
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        page  query     int  false  "Page pagination"
// @Success      200   {object}  []models.PostResponse
// @Success      209   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /posts [GET]
func GetPosts(c *gin.Context) {

	page := c.DefaultQuery("page", "1")

	pageNum, err := strconv.Atoi(page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "page query argument most be an int",
		})
		return
	}

	var postErr *utils.CustomError

	posts, postErr := services.GetPostsService(pageNum, "")

	if postErr != nil {
		c.JSON(postErr.Code, gin.H{
			"error": postErr.Message,
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}


// GetUsersPosts godoc
// @Summary      Fetches all user's posts in database
// @Description  Fetches all user's posts in database but with limit
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        page  		query     int  		false  "Page pagination"
// @param		 username	path	  string	true   "Username"
// @Success      200   		{object}  []models.PostResponse
// @Success      204   		{object}  map[string]interface{}
// @Failure      500   		{object}  map[string]interface{}
// @Router       /users/posts/{username} [GET]
func GetUsersPost(c *gin.Context){
	page := c.DefaultQuery("page", "1")

	pageNum, err := strconv.Atoi(page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "page query argument most be an int",
		})
		return
	}

	username := c.Param("username")

	var postErr *utils.CustomError

	posts, postErr := services.GetPostsService(pageNum, username)

	if postErr != nil {

		c.JSON(postErr.Code, gin.H{
			"error": postErr.Message,
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// GetPost godoc
// @Summary      Fetches a specific post in database
// @Description  Fetches a specific post in database
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        id    path     int  true  "Post id"
// @Success      200   {object}  []models.PostResponse
// @Success      404   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /posts/{id} [GET]
func GetPost(c *gin.Context) {
	
	postId := c.Param("id")


	postIdInt, err := strconv.Atoi(postId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id, route parameter most be an int",
		})
		return
	}

	post, postErr := services.GetPost(postIdInt)

	if postErr != nil {
		c.JSON(postErr.Code, gin.H{
			"error": postErr.Message,
		})
		return
	}

	c.JSON(http.StatusOK, post)
}


// NewPost godoc
// @Summary      Write a new post
// @Description  Write a new post with title and body
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        post  body      models.PostRequest  true  "Post data"
// @Success      201   {object}  models.PostResponse
// @Failure      400   {object}  map[string]interface{}
// @Failure      404   {object}  map[string]interface{}
// @Router       /posts/write [post]
func NewPost(c *gin.Context) {

	reqBody := &models.PostRequest{} 

	username, ok := c.Get("user")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "authorization header missed, make sure you have a valid token in headers",
		})
		return
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		
		allErrors := utils.GenerateUserFriendlyError(err)


		if allErrors == nil{
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

	result, err := services.CreatePostService(reqBody, username.(string))


	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}


	post := models.MakePostResponse(result)

	c.JSON(http.StatusCreated, post)

}