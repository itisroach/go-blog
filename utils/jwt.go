package utils

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/itisroach/go-blog/models"
	"github.com/itisroach/go-blog/repositories"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWTToken(username string, isRefresh bool) (string, error) {
	

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
			"exp": time.Now().Add(time.Minute * 15).Unix(),
			"type": "access",
		}
	}
	
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	if isRefresh{
		repositories.SaveTokenToDB(tokenString, claim["username"].(string), time.Unix(claim["exp"].(int64), 0))
	}
	
	return tokenString, nil
}



func VerifyToken(token string) (*models.TokenClaims, *CustomError) {

	tokenParsed, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})


	if err != nil {

		return nil, &CustomError{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		}
	}


	if !tokenParsed.Valid {
		return nil, &CustomError{
			Code: http.StatusBadRequest,
			Message: "invalid token",
		}
	}

	claims, ok := tokenParsed.Claims.(*jwt.MapClaims)


	if !ok {
		return nil, &CustomError{
			Code: http.StatusBadRequest,
			Message: "invalid token claims",
		}
	}
	username, _ := (*claims)["username"].(string)
	tokenType, _ := (*claims)["type"].(string)
	expFloat, ok := (*claims)["exp"].(float64)

	if !ok {
		return nil, &CustomError{
			Code:    http.StatusBadRequest,
			Message: "invalid expiration claim",
		}
	}

	expTime := time.Unix(int64(expFloat), 0)

	if time.Now().After(expTime) {
		return nil, &CustomError{
			Code:    http.StatusUnauthorized,
			Message: "token has expired",
		}
	}


	return &models.TokenClaims{
		Username: username,
		Type:     tokenType,
		ExpTime:  expTime,
	}, nil

}