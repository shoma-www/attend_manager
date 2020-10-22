package service

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/shoma-www/attend_manager/api/entity"
)

// Session interface
type Session interface {
	Start(ctx context.Context, values entity.Store) (*entity.Session, error)
	Get(ctx context.Context, id entity.SessionID) (*entity.Session, error)
	Destoroy(ctx context.Context, s entity.Session) error
}

type RedisSession struct {
	cl *redis.Client
}

func NewRedisSession(cl *redis.Client) Session {
	return &RedisSession{
		cl: cl,
	}
}

func (RedisSession) Start(ctx context.Context, values entity.Store) (*entity.Session, error) {
	return nil, nil
}

func (RedisSession) Get(ctx context.Context, id entity.SessionID) (*entity.Session, error) {
	return nil, nil
}

func (RedisSession) Destoroy(ctx context.Context, s entity.Session) error {
	return nil
}
