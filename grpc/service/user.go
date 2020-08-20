package service

import (
	"context"
	"errors"

	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/entity"
)

// User user service
type User struct {
	logger core.Logger
	ur     UserRepository
}

// NewUser user service constructor
func NewUser(l core.Logger, ur UserRepository) *User {
	return &User{
		logger: l,
		ur:     ur,
	}
}

// Register ユーザ登録
func (u *User) Register(ctx context.Context, userID string, password string) error {
	u.logger.WithUUID(ctx).Debug("register user_id: %s", userID)
	if u, err := u.ur.Get(userID, password); err != entity.ErrUserNotFound {
		if u == nil {
			err = errors.New("cannot use user")
		}
		return err
	}
	if err := u.ur.Register(userID, password); err != nil {
		return err
	}
	return nil
}

// UserRepository Access Interface
type UserRepository interface {
	Get(userID string, password string) (*entity.User, error)
	Register(userID string, password string) error
}
