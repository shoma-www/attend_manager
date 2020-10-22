package service

import (
	"context"

	"github.com/shoma-www/attend_manager/api/entity"
	"github.com/shoma-www/attend_manager/core"
)

// User operation
type User struct {
	ur UserRepository
}

// NewUser コンストラクタ
func NewUser(ur UserRepository) *User {
	return &User{ur: ur}
}

// Register ユーザーの登録
func (u *User) Register(ctx context.Context, user entity.User) error {
	l := core.GetLogger(ctx)
	l.Info("Create Login ID: %s", user.LoginID)
	if err := u.ur.Resister(ctx, user); err != nil {
		l.Error("Failed Create User")
		return err
	}
	l.Info("Success Create User")
	return nil
}

// SigIn 認証するで
func (u *User) SigIn(ctx context.Context, groupName, loginID, password string) error {
	l := core.GetLogger(ctx)
	l.Info("SignIn Login ID: %s", loginID)
	if err := u.ur.SigIn(ctx, groupName, loginID, password); err != nil {
		l.Error("Failed SinIn User")
		return err
	}
	l.Info("Success SinIn")
	return nil
}
