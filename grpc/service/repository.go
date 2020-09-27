package service

import (
	"context"

	"github.com/rs/xid"
	"github.com/shoma-www/attend_manager/grpc/entity"
)

// Transaction wrapper
type Transaction interface {
	Transaction(ctx context.Context, target func(tctx context.Context) (interface{}, error)) (interface{}, error)
}

// UserRepository Access Interface
type UserRepository interface {
	Get(ctx context.Context, groupID xid.ID, loginID string) (*entity.User, error)
	Register(ctx context.Context, groupID xid.ID, loginID, password, name string) (*entity.User, error)
}
