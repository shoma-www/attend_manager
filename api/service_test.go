package main

import (
	"context"
	"testing"

	"github.com/shoma-www/attend_manager/core"
)

type dummyCheckRepository struct {
	hs  *HealthCheckStatus
	err error
}

func (dc *dummyCheckRepository) HealthCheck(ctx context.Context) (*HealthCheckStatus, error) {
	return dc.hs, dc.err
}

func TestCheckService_HealthCheck(t *testing.T) {
	l := core.NewLogger(core.Debug)
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		rep := &dummyCheckRepository{
			hs:  &HealthCheckStatus{Status: "success"},
			err: nil,
		}
		hs := CheckService{
			logger: l,
			repo:   rep,
		}
		if err := hs.HealthCheck(ctx); err != nil {
			t.Errorf("CheckService.HealthCheck() error = %v, wantErr nil", err)
		}
	})
}
