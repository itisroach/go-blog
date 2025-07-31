package utils

import "golang.org/x/crypto/bcrypt"

func HashString(value string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)


	if err != nil {
		return "", err
	}


	return string(hashedPassword), nil

}


func ComparePassword(value string, hashed string) (bool, error) {
	
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(value))

	if err != nil {
		return false, err
	}

	return true, nil
}