package mysql

import (
	"hireplus-project/internal/entity"
	"hireplus-project/internal/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByPhone(phone string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("phone_number = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(userID string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *entity.User) error {
	return r.db.Save(user).Error
}
