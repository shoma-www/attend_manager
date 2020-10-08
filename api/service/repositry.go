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
}
