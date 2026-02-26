package di

import "memoflow/internal/user"

type IStatRepository interface {
	AddClick(memoId uint)
}

type IUserRepository interface {
	Create(user *user.User) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
}
