package repositories

import (
	"errors"

	"github.com/itisroach/go-blog/database"
	"github.com/itisroach/go-blog/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		Preload("User").
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



func UpdatePost(payload models.UpdatePostRequest ,id int, username string) (*models.PostResponse, error) {

	var post models.Post

	updates := make(map[string]interface{})

	if payload.Title != "" {
		updates["title"] = payload.Title
	}

	if payload.Body != "" {
		updates["body"] = payload.Body
	}

	if len(updates) == 0 {
		return nil, errors.New("no fields specified to update")
	}


	err := database.DB.
		Model(&post).
		Clauses(
			clause.From{Tables: []clause.Table{{Name: "users"}}},                          
		).
		Where("posts.id = ? AND posts.user_id = users.id AND users.username = ?", id, username).
		Updates(updates).Error 


	if err != nil {
		return nil, err
	}


	err = database.DB.
		Joins("JOIN users ON users.id = posts.user_id").
		Where("posts.id = ? AND users.username = ?", id, username).
		Preload("User").
		First(&post).Error

	if err != nil {
		return nil, err
	}

	return models.MakePostResponse(&post), nil

}



func DeletePost(id int, username string) error {

	result := database.DB.
	Clauses(
		clause.From{Tables: []clause.Table{{Name: "users"}}},                          
	).
	Where("posts.id = ? AND posts.user_id = users.id AND users.username = ?", id, username).
	Delete(&models.Post{})

	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}


	return nil
}