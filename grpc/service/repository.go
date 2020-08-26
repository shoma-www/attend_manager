package service

import (
	"context"

	"github.com/shoma-www/attend_manager/grpc/entity"
)

// Transaction wrapper
type Transaction interface {
	Transaction(ctx context.Context, target func(tctx context.Context) (interface{}, error)) (interface{}, error)
}

// UserRepository Access Interface
type UserRepository interface {
	Get(ctx context.Context, userID string) ([]*entity.User, error)
	Register(ctx context.Context, userID string, password string) (*entity.User, error)
}
