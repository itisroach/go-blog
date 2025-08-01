package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primarykey"`
	Username string `gorm:"uniqueindex:idx_username"`
	Name     string	`gorm:"default:Unknown"`
	Password string `gorm:"<-:create"`
}


type UserRequest struct {
	Username 	string `binding:"required,min=4,max=20"`
	Name 		string 
	Password	string `binding:"required,min=8,max=64"`
}


type UserResponse struct {
	Id 			int
	Username	string
	Name 		string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}


func NewUserResponse(u *User) *UserResponse{
	return &UserResponse{
		Id:       u.ID,
		Username: u.Username,
		Name:     u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}


func (u *UserRequest) MakeUser() *User {
	return &User{
		Name: u.Name,
		Username: u.Username,
		Password: u.Password,
	}
}



type LoginRequest struct {
	Username 	string	`binding:"required,min=4,max=20"`
	Password	string	`binding:"required,min=8,max=64"`
}