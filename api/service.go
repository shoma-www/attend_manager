package main

import (
	"context"

	"github.com/shoma-www/attend_manager/api/entity"
	"github.com/shoma-www/attend_manager/core"
)

// CheckService 他システムへのhealthcheck
type CheckService struct {
	logger core.Logger
	repo   CheckRepository
}

// NewCheckService constructor
func NewCheckService(l core.Logger, cr CheckRepository) CheckService {
	return CheckService{logger: l, repo: cr}
}

// HealthCheck GrpcのサーバーにHealthCheckを実施
func (hs CheckService) HealthCheck(ctx context.Context) error {
	status, err := hs.repo.HealthCheck(ctx)
	if err != nil {
		return err
	}
	hs.logger.Debug("health check grpc status: %s", status.Status)
	return err
}

// CheckRepository Repository
type CheckRepository interface {
	HealthCheck(ctx context.Context) (*entity.HealthCheckStatus, error)
}
