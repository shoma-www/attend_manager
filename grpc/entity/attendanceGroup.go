package entity

import (
	"errors"

	"github.com/rs/xid"
)

// AttendanceGroup type
type AttendanceGroup struct {
	ID       xid.ID
	LoginID  string
	Password string
	Name     string
}

// Error UserのError定義
var (
	ErrAttendanceGroupNotFound   = errors.New("attendance group not found")
	ErrDuplicatedAttendanceGroup = errors.New("group is duplicated")
)
