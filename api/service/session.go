package service

import (
	"context"

	"github.com/shoma-www/attend_manager/api/entity"
)

// Session interface
type Session interface {
	Start(ctx context.Context, values entity.Store) (*entity.Session, error)
	Get(ctx context.Context, id entity.SessionID) (*entity.Session, error)
	Destoroy(ctx context.Context, s entity.Session) error
}

// const lifetime = 60 * time.Hour

type RedisSession struct {
	addr     string
	password string
	db       int
}

func NewRedisSession(addr, password string, db int) Session {
	return &RedisSession{
		addr:     addr,
		password: password,
		db:       db,
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
