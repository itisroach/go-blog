package repositories

import (
	"errors"

	"github.com/itisroach/go-blog/database"
	"github.com/itisroach/go-blog/models"
)

func CreatePost(post *models.Post) error {


	result := database.DB.Create(&post)


	if result.Error != nil {
		return result.Error
	}

	return nil

}


func GetPosts(page int) (*[]models.PostResponse ,error) {

	pageSize := 10

	offset := (page - 1) * pageSize
	
	var posts *[]models.Post
	
	err := database.DB.Preload("User").Limit(pageSize).Offset(offset).Find(&posts).Error

	if err != nil {
		return nil, errors.New("failed to fetch posts")
	}

	var result []models.PostResponse

	for _, post := range *posts {
		result = append(result, *models.MakePostResponse(&post))
	}

	return &result, nil
}