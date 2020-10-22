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
	SigIn(ctx context.Context, groupName, loginID, password string) error
}

// GroupRepository Repository
type GroupRepository interface {
	Create(ctx context.Context, group entity.Group, user entity.User) error
}

// SessionRepository Repository
type SessionRepository interface {
	Start(ctx context.Context, values entity.Store) (entity.Session, error)
	Get(ctx context.Context, id entity.SessionID) (entity.Session, error)
	Destoroy(ctx context.Context, s entity.Session) error
}
