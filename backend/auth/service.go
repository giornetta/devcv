package auth

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// Service will handle authentication for the API
type Service interface {
	GenerateToken(username string) (string, error)
	Authenticate(bearer, username string) error
}

type service struct {
	Key string
}

// New returns an Authenticator Service
func New(key string) Service {
	return &service{
		Key: key,
	}
}

// GenerateToken generates a JWT based on a username
func (s *service) GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	return token.SignedString([]byte(s.Key))
}

// Authenticate will check if the given token matches the user that should be calling a protected method
func (s *service) Authenticate(bearer, username string) error {
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		bearer = bearer[7:]
	}

	token, err := jwt.Parse(bearer, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(s.Key), nil
	})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return errors.New("invalid token")
	}

	if claims["username"] != username {
		return errors.New("invalid token")
	}

	return nil
}
