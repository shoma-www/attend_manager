package entity

import (
	"errors"

	"github.com/rs/xid"
)

// User User type
type User struct {
	ID       xid.ID
	GroupID  xid.ID
	LoginID  string
	Password string
	Name     string
}

// Error UserのError定義
var (
	ErrUserNotFound   = errors.New("user not found")
	ErrDuplicatedUser = errors.New("user id is duplicated")
)
