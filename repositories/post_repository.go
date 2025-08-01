package repositories

import (
	"errors"

	"github.com/itisroach/go-blog/database"
	"github.com/itisroach/go-blog/models"
	"gorm.io/gorm"
)

func CreatePost(post *models.Post) error {


	result := database.DB.Create(&post)


	if result.Error != nil {
		return result.Error
	}

	return nil

}


func GetPosts(page int, username string) (*[]models.PostResponse ,error) {

	pageSize := 10

	offset := (page - 1) * pageSize
	
	var posts []models.Post
	
	var err error;

	switch username {
	case "":
		err = database.DB.Preload("User").Limit(pageSize).Offset(offset).Find(&posts).Error
	
	default:
		err = database.DB.
		Joins("JOIN users ON users.id = posts.user_id").
    	Where("users.username = ?", username).
		Limit(pageSize).
		Offset(offset).
		Find(&posts).Error
	}

	if username != "" && len(posts) == 0 {
		return nil, errors.New("user does not have any posts")
	}
	
	if len(posts) == 0 {
		return nil, errors.New("we don't have any posts for now")
	}
	
	if err != nil {
		return nil, errors.New("failed to fetch posts")
	}

	var result []models.PostResponse

	for _, post := range posts {
		result = append(result, *models.MakePostResponse(&post))
	}

	return &result, nil
}


func GetSinglePost(id int) (*models.PostResponse, error) {

	var post models.Post

	err := database.DB.Preload("User").Where("id = ?", id).First(&post).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("post not found")
	}

	if err != nil {
		return nil, errors.New("failed to fetch posts")
	}

	return models.MakePostResponse(&post), nil

}