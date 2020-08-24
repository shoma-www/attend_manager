package service

import (
	"context"

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
	u.ur.Transaction(ctx, func(tctx context.Context) error {
		us, err := u.ur.Get(tctx, userID)
		if err != entity.ErrUserNotFound {
			return err
		}
		if len(us) > 0 {
			return entity.ErrDuplicatedUser
		}
		hashedPassword, err := core.GenerateHashedPassword(password)
		if err != nil {
			return err
		}
		user, err := u.ur.Register(tctx, userID, string(hashedPassword))
		if err != nil {
			return err
		}
		u.logger.WithUUID(tctx).Debug("create user: %v", user)
		return nil
	})

	return nil
}
