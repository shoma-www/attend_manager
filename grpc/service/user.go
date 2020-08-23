package service

import (
	"context"

	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/entity"
	"golang.org/x/crypto/bcrypt"
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
	us, err := u.ur.Get(ctx, userID)
	u.logger.WithUUID(ctx).Debug("%v", us)
	if err != entity.ErrUserNotFound {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.logger.WithUUID(ctx).Debug("hash: %s", hashedPassword)
	user, err := u.ur.Register(ctx, userID, string(hashedPassword))
	if err != nil {
		return err
	}
	u.logger.WithUUID(ctx).Debug("create user: %v", user)
	return nil
}

// UserRepository Access Interface
type UserRepository interface {
	Get(ctx context.Context, password string) ([]*entity.User, error)
	Register(ctx context.Context, userID string, password string) (*entity.User, error)
}
