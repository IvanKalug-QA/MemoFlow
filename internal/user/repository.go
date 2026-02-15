package user

import (
	"memoflow/pkg/db"
)

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{Database: database}
}

type UserRepository struct {
	Database *db.Db
}

func (u *UserRepository) Create(user *User) (*User, error) {
	result := u.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := u.Database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
