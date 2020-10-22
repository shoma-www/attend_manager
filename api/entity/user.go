package entity

import "errors"

// User User type
type User struct {
	ID       string
	GroupID  string
	LoginID  string
	Password string
	Name     string
}

var (
	ErrUnauthenticated = errors.New("not authenticated")
)
