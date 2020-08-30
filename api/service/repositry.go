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
	Resister(ctx context.Context, userID string, password string) error
}
