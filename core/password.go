package core

import "golang.org/x/crypto/bcrypt"

// GenerateHashedPassword is generated from password
func GenerateHashedPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CompareHashAndPassword compare a hashed password and a password
func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
