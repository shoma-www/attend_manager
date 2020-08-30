package service

import (
	"context"
	"errors"

	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/entity"
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
func (u *User) Register(ctx context.Context, userID string, password string) (*entity.User, error) {
	return nil, errors.New("test")
}
