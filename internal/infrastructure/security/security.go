package security

import (
	"backend-kata/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type Security struct {
	SecurityConfig config.SecurityConfig
}

func NewSecurity(securityConfig config.SecurityConfig) Security {
	return Security{SecurityConfig: securityConfig}
}

func (s Security) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s Security) CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
func (s Security) GenerateJWT(data map[string]string) (*string, error) {
	claims := jwt.MapClaims{}
	for key, value := range data {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(s.SecurityConfig.JWTSecret)

	tokenString, errorSignedToken := token.SignedString(secretKey)
	if errorSignedToken != nil {
		log.Info().Msgf("Error signing the token: %v", errorSignedToken)
		return nil, errorSignedToken
	}

	return &tokenString, nil
}

func (s Security) IsValidateJWT(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.SecurityConfig.JWTSecret, nil
	})

	if err != nil || !token.Valid {
		return false
	}

	return true
}
