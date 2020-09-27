package service

import (
	"context"

	"github.com/rs/xid"
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
func (u *User) Register(ctx context.Context, groupID xid.ID, loginID, password, name string) (*entity.User, error) {
	u.logger.WithUUID(ctx).Info("register user. id: %s", loginID)
	v, err := u.tr.Transaction(ctx, func(tctx context.Context) (interface{}, error) {
		if us, err := u.ur.Get(tctx, groupID, loginID); err != entity.ErrUserNotFound {
			if us != nil {
				return nil, entity.ErrDuplicatedUser
			}
			return nil, err
		}

		hashedPassword, err := core.GenerateHashedPassword(password)
		if err != nil {
			return nil, err
		}
		user, err := u.ur.Register(tctx, groupID, loginID, string(hashedPassword), name)
		if err != nil {
			return nil, err
		}
		u.logger.WithUUID(tctx).Info("complete registered user. login id: %s, id: %s", user.LoginID, user.ID)
		return user, nil
	})

	if user, ok := v.(*entity.User); ok {
		return user, err
	}
	return nil, err
}
