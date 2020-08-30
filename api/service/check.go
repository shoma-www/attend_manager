package service

import (
	"context"

	"github.com/shoma-www/attend_manager/core"
)

// Check 他システムへのhealthcheck
type Check struct {
	logger  core.Logger
	factory CheckRepository
}

// NewCheck constructor
func NewCheck(l core.Logger, cr CheckRepository) Check {
	return Check{logger: l, factory: cr}
}

// HealthCheck GrpcのサーバーにHealthCheckを実施
func (hs Check) HealthCheck(ctx context.Context) error {
	status, err := hs.factory.HealthCheck(ctx)
	if err != nil {
		return err
	}
	hs.logger.Debug("health check grpc status: %s", status.Status)
	return err
}
