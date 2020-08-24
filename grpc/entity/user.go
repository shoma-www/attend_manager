package entity

import "errors"

// User User type
type User struct {
	ID       string
	UserID   string
	Password string
}

// Error UserのError定義
var (
	ErrUserNotFound   = errors.New("user not found")
	ErrDuplicatedUser = errors.New("user id is duplicated")
)
