package service

import (
	"context"

	"github.com/shoma-www/attend_manager/api/entity"
)

// CheckRepository Repository
type CheckRepository interface {
	HealthCheck(ctx context.Context) (*entity.HealthCheckStatus, error)
}

// UserRepository Repository
type UserRepository interface {
	Resister(ctx context.Context, user entity.User) error
	SigIn(ctx context.Context, groupName, loginID, password string) (*entity.SigninData, error)
}

// GroupRepository Repository
type GroupRepository interface {
	Create(ctx context.Context, group entity.Group, user entity.User) error
}
