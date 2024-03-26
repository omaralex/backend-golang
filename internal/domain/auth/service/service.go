package service

import (
	"backend-kata/internal/domain/auth"
	customError "backend-kata/internal/domain/errors"
	"context"
	"fmt"
)

type Repository interface {
	SaveUser(user auth.User) error
	GetUserByEmail(email string) (*auth.User, error)
}

type Security interface {
	HashPassword(password string) (string, error)
	CheckPassword(password, hashedPassword string) bool
	GenerateJWT(data map[string]string) (*string, error)
}
type AuthService struct {
	repository Repository
	Security   Security
}

func NewAuthService(repository Repository, security Security) AuthService {
	return AuthService{repository, security}
}
func (s AuthService) SignUp(ctx context.Context, user auth.User) error {
	hashedPassword, errorHashPassword := s.Security.HashPassword(user.Password)
	if errorHashPassword != nil {
		return fmt.Errorf("error hashing password: %v", errorHashPassword)
	}

	user.Password = hashedPassword

	errorSaveUser := s.repository.SaveUser(user)
	if errorSaveUser != nil {
		return errorSaveUser
	}

	return nil
}
func (s AuthService) Authenticate(ctx context.Context, email string, password string) (*string, error) {
	user, errorGetUser := s.repository.GetUserByEmail(email)
	if errorGetUser != nil {
		return nil, errorGetUser
	}

	if user != nil {
		return nil, &customError.UserNotFoundError{}
	}

	if !s.Security.CheckPassword(password, user.Password) {
		return nil, &customError.InvalidPasswordError{}
	}

	dataJWT := map[string]string{
		"email": user.Email,
	}

	jwtGenerated, errorGenerateJWT := s.Security.GenerateJWT(dataJWT)
	if errorGenerateJWT != nil {
		return nil, errorGenerateJWT
	}

	return jwtGenerated, nil
}
