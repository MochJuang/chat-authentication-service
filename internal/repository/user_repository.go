package repository

import "hireplus-project/internal/entity"

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByPhone(phone string) (*entity.User, error)
	GetUserByID(userID string) (*entity.User, error)
	UpdateUser(user *entity.User) error
}
