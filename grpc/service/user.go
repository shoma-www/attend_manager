package service

import (
	"context"

	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/entity"
)

// User user service
type User struct {
	logger core.Logger
	tr     Transaction
	ur     UserRepository
}

// NewUser user service constructor
func NewUser(l core.Logger, tr Transaction, ur UserRepository) *User {
	return &User{
		logger: l,
		tr:     tr,
		ur:     ur,
	}
}

// Register ユーザ登録
func (u *User) Register(ctx context.Context, userID string, password string) (*entity.User, error) {
	u.logger.WithUUID(ctx).Info("register user. id: %s", userID)
	v, err := u.tr.Transaction(ctx, func(tctx context.Context) (interface{}, error) {
		if us, err := u.ur.Get(tctx, userID); err != entity.ErrUserNotFound {
			if us != nil && len(us) > 0 {
				return nil, entity.ErrDuplicatedUser
			}
			return nil, err
		}

		hashedPassword, err := core.GenerateHashedPassword(password)
		if err != nil {
			return nil, err
		}
		user, err := u.ur.Register(tctx, userID, string(hashedPassword))
		if err != nil {
			return nil, err
		}
		u.logger.WithUUID(tctx).Info("complete registered user. id: %s, uuid: %s", user.UserID, user.ID)
		return user, nil
	})

	if user, ok := v.(*entity.User); ok {
		return user, err
	}
	return nil, err
}
