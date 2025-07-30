package models



type User struct {
	ID 			int		`gorm:"primarykey"`
	Username	string	`gorm:"uniqueindex:idx_username"`
	Name		string
	Password 	string	
}