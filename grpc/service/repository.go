package service

import (
	"context"

	"github.com/shoma-www/attend_manager/grpc/entity"
)

// Transaction wrapper
type Transaction interface {
	Transaction(ctx context.Context, target func(tctx context.Context) error) error
}

// UserRepository Access Interface
type UserRepository interface {
	Transaction
	Get(ctx context.Context, password string) ([]*entity.User, error)
	Register(ctx context.Context, userID string, password string) (*entity.User, error)
}
