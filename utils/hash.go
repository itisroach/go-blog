package utils

import "golang.org/x/crypto/bcrypt"

func HashString(value string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)


	if err != nil {
		return "", err
	}


	return string(hashedPassword), nil

}