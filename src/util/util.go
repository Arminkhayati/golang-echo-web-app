package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "",err
	}
	return string(bytes), nil
}

func ComparePassWithHash(normal string, hashed string)(ok bool){
	err := bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(normal))
	if err != nil {
		ok = false
		return
	}
	ok = true
	return
}