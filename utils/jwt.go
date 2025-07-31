package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(username string, isRefresh bool) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	var claim jwt.MapClaims

	switch isRefresh {
	case true:
		claim = jwt.MapClaims{
			"username": username,
			"exp": time.Now().Add(time.Hour * 48).Unix(),
			"type": "refresh",
		}
	default:
		claim = jwt.MapClaims{
			"username": username,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
			"type": "access",
		}
	}
	
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}