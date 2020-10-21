package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/entity"
)

var ErrUnauthorized = errors.New("unauthorized group id or login id or password")

// User user service
type User struct {
	logger core.Logger
	tr     Transaction
	ur     UserRepository
	ag     AttendanceGroupRepository
}

// NewUser user service constructor
func NewUser(l core.Logger, tr Transaction, ur UserRepository, ag AttendanceGroupRepository) *User {
	return &User{
		logger: l,
		tr:     tr,
		ur:     ur,
		ag:     ag,
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

// SignIn サインイン
func (u *User) SignIn(ctx context.Context, groupName, loginID, password string) (string, error) {
	l := core.GetLogger(ctx)
	l.Info("start signin user group Name: %s, login ID: %s", groupName, loginID)
	g, err := u.ag.Get(ctx, groupName)
	if err != nil {
		if err == entity.ErrAttendanceGroupNotFound {
			l.Warn(errors.Wrapf(err, "not found group. group name: %s", groupName).Error())
			return "", ErrUnauthorized
		}
		return "", err
	}
	u1, err := u.ur.Get(ctx, g.ID, loginID)
	if err != nil {
		if err == entity.ErrUserNotFound {
			l.Warn(errors.Wrapf(err, "not found user. login id: %s", loginID).Error())
			return "", ErrUnauthorized
		}
		return "", err
	}
	if err = core.CompareHashAndPassword(u1.Password, password); err != nil {
		l.Warn(errors.Wrap(err, "invalid password").Error())
		return "", ErrUnauthorized
	}
	l.Info("success signin user group Name: %s, login ID: %s", groupName, loginID)
	return u1.Name, err
}
