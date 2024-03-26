package dtos

import (
	"backend-kata/internal/domain/auth"
	"github.com/google/uuid"
)

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request *SignUpRequest) ToUser() auth.User {
	return auth.User{
		ID:       uuid.New().ID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}
