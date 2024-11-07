package repository

import (
	"go-rest-api/model"
)

type IUserRepository interface {
	CreateUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}
