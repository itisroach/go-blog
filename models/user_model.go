package models

import (
	"fmt"
	"time"

	"github.com/itisroach/go-blog/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primarykey"`
	Username string `gorm:"uniqueindex:idx_username" binding:"required,min=4,max=20"`
	Name     string	`gorm:"default:Unknown"`
	Password string `gorm:"<-:create" binding:"required,min=8,max=64"`
}


type UserResponse struct {
	Id 			int
	Username	string
	Name 		string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}


func (u *User) NewUserResponse() *UserResponse{
	return &UserResponse{
		Id:       u.ID,
		Username: u.Username,
		Name:     u.Name,
	}
}


func (u *User) HashPassword() error {
	var err error
	u.Password, err = utils.HashString(u.Password)

	if err != nil {
		return err
	}
	fmt.Println(*u)
	return nil
}