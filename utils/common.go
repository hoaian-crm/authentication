package utils

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func RandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func ComparePassword(password string, cipherText string) bool {
	if password == "" {
		return false
	}
	ok := bcrypt.CompareHashAndPassword([]byte(password), []byte(cipherText))
	return ok == nil
}
