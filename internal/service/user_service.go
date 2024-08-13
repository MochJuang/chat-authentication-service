package service

import (
	"fmt"
	"notification-service/internal/config"
	"notification-service/internal/entity"
	e "notification-service/internal/exception"
	"notification-service/internal/model"
	"notification-service/internal/repository"
	"notification-service/internal/utils"
)

type UserService interface {
	Register(request model.UserRegisterRequest) (*model.UserResponse, error)
	GetUserByID(userID int) (*entity.User, error)
}

type userService struct {
	userRepo repository.UserRepository
	config   config.Config
}

func NewUserService(userRepo repository.UserRepository, cfg config.Config) UserService {
	return &userService{userRepo: userRepo, config: cfg}
}

func (s *userService) Register(request model.UserRegisterRequest) (*model.UserResponse, error) {
	var err error
	err = utils.Validate(request)
	if err != nil {
		return nil, err
	}

	password, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, e.Internal(err)
	}

	var user = entity.User{
		Username: request.Username,
		Email:    request.Email,
		Password: password,
	}

	err = s.userRepo.CreateUser(&user)
	if err != nil {
		return nil, e.Internal(err)
	}

	response := model.ToUserResponse(user)

	return &response, nil
}

func (s *userService) GetUserByID(userID int) (*entity.User, error) {
	if userID == 0 {
		return nil, e.Validation(fmt.Errorf("userID must be greater than 0"))
	}

	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, e.NotFound("user not found")
	}

	return user, nil
}
