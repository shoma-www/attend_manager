package infra

import "github.com/shoma-www/attend_manager/grpc/entity"

type userDAO struct {
}

func (u *userDAO) Get(userID string, password string) (*entity.User, error) {
	return nil, entity.ErrUserNotFound
}

func (u *userDAO) Register(userID string, password string) error {
	return nil
}
