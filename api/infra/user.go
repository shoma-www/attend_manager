package infra

import (
	"context"
	"errors"

	"github.com/shoma-www/attend_manager/api/entity"
)

type userGrpc struct {
	address string
}

func (ug *userGrpc) Resister(ctx context.Context, userID string, password string) (*entity.User, error) {
	return nil, errors.New("register error")
}
