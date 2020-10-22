package service

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/hex"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/securecookie"
	"github.com/shoma-www/attend_manager/api/entity"
)

// Session interface
type Session interface {
	Start(ctx context.Context, values entity.Store) (*entity.Session, error)
	Get(ctx context.Context, id entity.SessionID) (*entity.Session, error)
	Destoroy(ctx context.Context, s entity.Session) error
}

const lifetime = 60 * time.Hour

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

func (rs RedisSession) Start(ctx context.Context, values entity.Store) (*entity.Session, error) {
	client := rs.newClient()
	defer client.Close()

	key := hex.EncodeToString(securecookie.GenerateRandomKey(64))
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	if err := e.Encode(values); err != nil {
		return nil, err
	}

	if err := client.Set(key, b.Bytes(), lifetime).Err(); err != nil {
		return nil, err
	}
	return &entity.Session{
		ID:     entity.SessionID(key),
		Values: values,
	}, nil
}

func (rs RedisSession) Get(ctx context.Context, id entity.SessionID) (*entity.Session, error) {
	client := rs.newClient()
	defer client.Close()

	b, err := client.Get(string(id)).Result()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer([]byte(b))
	values := make(entity.Store)
	d := gob.NewDecoder(buf)
	if err := d.Decode(&values); err != nil {
		return nil, err
	}
	return &entity.Session{
		ID:     id,
		Values: values,
	}, nil
}

func (rs RedisSession) Destoroy(ctx context.Context, s entity.Session) error {
	client := rs.newClient()
	defer client.Close()
	if err := client.Del(string(s.ID)).Err(); err != nil {
		return err
	}
	return nil
}

func (rs RedisSession) newClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     rs.addr,
		Password: rs.password,
		DB:       rs.db,
	})
}
