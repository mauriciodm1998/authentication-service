package token

import (
	"authentication-service/internal/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId string, email string) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId
	permissions["email"] = email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.Get().Token.Key))
}
