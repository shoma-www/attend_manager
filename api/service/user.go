package service

import (
	"context"

	"github.com/shoma-www/attend_manager/core"
)

// User operation
type User struct {
	logger core.Logger
	ur     UserRepository
}

// NewUser コンストラクタ
func NewUser(l core.Logger, ur UserRepository) *User {
	return &User{logger: l, ur: ur}
}

// Register ユーザーの登録
func (u *User) Register(ctx context.Context, userID string, password string) error {
	u.logger.WithUUID(ctx).Info("Create User ID: %s", userID)
	err := u.ur.Resister(ctx, userID, password)
	u.logger.WithUUID(ctx).Info("Success Create User")
	return err
}
