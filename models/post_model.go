package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID	uint 
	User  	User 		`gorm:"OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title 	string
	Body  	string
}


type PostRequest struct {
	Title		string	`binding:"required,min=8,max=128"`
	Body		string	`binding:"required,min=24"`
}


type PostResponse struct {
	ID 			uint
	User 		*UserResponse
	Title 		string
	Body		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}


func MakePostResponse(post *Post) *PostResponse {
	return &PostResponse{
		ID: post.ID,
		User: NewUserResponse(&post.User),
		Title: post.Title,
		Body: post.Body,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func (p *PostRequest) MakePost(user *User) *Post {
	return &Post{
		UserID: uint(user.ID),
		User: *user,
		Title: p.Title,
		Body: p.Body,
	}
}