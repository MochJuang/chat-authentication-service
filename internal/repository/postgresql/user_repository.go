package postgresql

import (
	"authentication-service/internal/entity"
	"authentication-service/internal/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
	tx map[string]*gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db, tx: make(map[string]*gorm.DB)}
}

func (r *userRepository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByUuid(uuid string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
