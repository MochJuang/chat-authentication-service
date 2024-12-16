package repository

import "authentication-service/internal/entity"

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByUuid(uuid string) (*entity.User, error)
	UpdateUser(user *entity.User) error
	GetUserByUsername(username string) (*entity.User, error)

	Begin() (string, error)
	Rollback(sessionId string) error
	Commit(sessionId string) error
}
