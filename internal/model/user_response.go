package model

import (
	"authentication-service/internal/entity"
	"time"
)

type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type AuthenticationResponse struct {
	AccessToken      string    `json:"access_token"`
	RefreshToken     string    `json:"refresh_token"`
	ExpiresIn        time.Time `json:"expires_in"`
	RefreshExpiresIn time.Time `json:"refresh_expires_in"`
}

func ToUserResponse(user entity.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
