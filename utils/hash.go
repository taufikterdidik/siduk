package utils

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
	"strconv"
)

func HashPIN(pin string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pin), 10)
	return string(bytes), err
}

func CheckPIN(pin, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pin)) == nil
}

func GeneratePIN() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(100000 + rand.Intn(900000)) // 6 digit
}