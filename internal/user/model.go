package user

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Email    string `gorm:"index"`
	Username string
	Password string
}
