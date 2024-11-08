package service

import (
	"authentication-service/internal/config"
	"authentication-service/internal/entity"
	e "authentication-service/internal/exception"
	"authentication-service/internal/model"
	"authentication-service/internal/repository"
	"authentication-service/internal/utils"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type UserService interface {
	Register(request model.UserRegisterRequest) (*model.AuthenticationResponse, error)
	GetUserByID(userID int) (*model.UserResponse, error)
	Login(request model.UserLoginRequest) (*model.AuthenticationResponse, error)
	RefreshToken(request model.RefreshTokenRequest) (*model.AuthenticationResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
	config   config.Config
}

func NewUserService(userRepo repository.UserRepository, cfg config.Config) UserService {
	return &userService{userRepo: userRepo, config: cfg}
}

func (s *userService) Register(request model.UserRegisterRequest) (*model.AuthenticationResponse, error) {
	err := utils.Validate(request)
	if err != nil {
		return nil, err
	}

	if _, err = s.userRepo.GetUserByUsername(request.Username); err == nil {
		return nil, e.BadRequest("username already exists")
	}

	password, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, e.Internal(err)
	}

	user := entity.User{
		UUID:     uuid.New().String(),
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: password,
	}

	err = s.userRepo.CreateUser(&user)
	if err != nil {
		return nil, e.Internal(err)
	}

	response, err := utils.AuthenticationResponse(user.UUID, s.config)
	if err != nil {
		return nil, e.Internal(err)
	}

	return response, nil
}

func (s *userService) Login(request model.UserLoginRequest) (*model.AuthenticationResponse, error) {
	err := utils.Validate(request)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetUserByUsername(request.Username)
	if err != nil {
		return nil, e.BadRequest("username or password is incorrect")
	}

	if !utils.CheckPasswordHash(request.Password, user.Password) {
		return nil, e.BadRequest("username or password is incorrect")
	}

	response, err := utils.AuthenticationResponse(user.UUID, s.config)
	if err != nil {
		return nil, e.Internal(err)
	}

	return response, nil
}

func (s *userService) RefreshToken(request model.RefreshTokenRequest) (*model.AuthenticationResponse, error) {
	err := utils.Validate(request)
	if err != nil {
		return nil, err
	}

	jwtKey := s.config.JWTSecret

	claims, err := utils.ParseToken(request.RefreshToken, jwtKey)
	if err != nil {
		return nil, e.BadRequest("invalid refresh token")
	}

	uuidUser := strings.Split(claims.UserID, "_")[0]
	fmt.Println("full", claims.UserID)
	fmt.Println("uuid", uuidUser)

	response, err := utils.AuthenticationResponse(uuidUser, s.config)
	if err != nil {
		return nil, e.Internal(err)
	}

	return response, nil
}

func (s *userService) GetUserByID(userID int) (*model.UserResponse, error) {
	if userID == 0 {
		return nil, e.BadRequest(fmt.Errorf("userID must be greater than 0"))
	}

	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, e.NotFound("user not found")
	}
	response := model.ToUserResponse(*user)

	return &response, nil
}
