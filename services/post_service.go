package services

import (
	"net/http"
	"strings"

	"github.com/itisroach/go-blog/models"
	"github.com/itisroach/go-blog/repositories"
	"github.com/itisroach/go-blog/utils"
)


func GetPostsService(page int, username string) (*[]models.PostResponse, *utils.CustomError) {

	posts, err := repositories.GetPosts(page, username)


	if err != nil {
		code := http.StatusInternalServerError

		if strings.Contains(err.Error(), "any posts") {
			code = http.StatusNoContent 
		}

		return nil, &utils.CustomError{
			Code: code,
			Message: err.Error(),
		}
	}


	return posts, nil

}


func GetPost(id int) (*models.PostResponse, *utils.CustomError) {
	post, err := repositories.GetSinglePost(id)

	if err != nil {
		code := http.StatusInternalServerError

		if strings.Contains(err.Error(), "found") {
			code = http.StatusNotFound
		}

		return nil, &utils.CustomError{
			Code: code,
			Message: err.Error(),
		}
	}

	return post, nil
}


func CreatePostService(reqBody *models.PostRequest) (*models.Post, *utils.CustomError) {
	
	user, err := repositories.GetUserRawData(reqBody.Username)	

	
	if err != nil {
		return nil, &utils.CustomError{
			Code: http.StatusNotFound,
			Message: "user not found",
		}
	}


	postInstanace := reqBody.MakePost(user)

	err = repositories.CreatePost(postInstanace)

	if err != nil {
		return nil, &utils.CustomError{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return postInstanace, nil
}