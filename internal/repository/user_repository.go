package repository

import "notification-service/internal/entity"

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByID(userID int) (*entity.User, error)
	UpdateUser(user *entity.User) error
}
