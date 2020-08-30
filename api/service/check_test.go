package service

import (
	"context"
	"testing"

	"github.com/shoma-www/attend_manager/api/entity"
	"github.com/shoma-www/attend_manager/core"
)

type dummyCheckRepository struct {
	hs  *entity.HealthCheckStatus
	err error
}

func (dc *dummyCheckRepository) HealthCheck(ctx context.Context) (*entity.HealthCheckStatus, error) {
	return dc.hs, dc.err
}

func TestCheck_HealthCheck(t *testing.T) {
	l := core.NewLogger(core.Debug)
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		rep := &dummyCheckRepository{
			hs:  &entity.HealthCheckStatus{Status: "success"},
			err: nil,
		}
		hs := Check{
			logger:  l,
			factory: rep,
		}
		if err := hs.HealthCheck(ctx); err != nil {
			t.Errorf("Check.HealthCheck() error = %v, wantErr nil", err)
		}
	})
}
