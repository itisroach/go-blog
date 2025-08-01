package services

import (
	"net/http"

	"github.com/itisroach/go-blog/models"
	"github.com/itisroach/go-blog/repositories"
	"github.com/itisroach/go-blog/utils"
)


func GetPostsService(page int) (*[]models.PostResponse, *utils.CustomError) {

	posts, err := repositories.GetPosts(page)


	if err != nil {
		return nil, &utils.CustomError{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}


	return posts, nil

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