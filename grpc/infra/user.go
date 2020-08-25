package infra

import (
	"context"

	"github.com/rs/xid"
	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/ent"
	"github.com/shoma-www/attend_manager/grpc/ent/user"
	"github.com/shoma-www/attend_manager/grpc/entity"
)

type userDAO struct {
	logger core.Logger
	cl     *ent.Client
}

func (ud *userDAO) Get(ctx context.Context, userID string) ([]*entity.User, error) {
	uc := ud.cl.User
	if tx, ok := getTX(ctx); ok {
		uc = tx.User
	}
	us, _ := uc.Query().Where(user.UserIDEQ(userID)).All(ctx)
	users := make([]*entity.User, 0, len(us))
	for _, u := range us {
		ud.logger.Debug("id: %s, uid: %s, pass: %s", u.UUID, u.UserID, u.Password)
		users = append(users, &entity.User{
			ID:       u.UUID.String(),
			UserID:   u.UserID,
			Password: u.Password,
		})
	}
	if len(users) > 0 {
		return users, nil
	}
	return nil, entity.ErrUserNotFound
}

func (ud *userDAO) Register(ctx context.Context, userID, password string) (*entity.User, error) {
	uc := ud.cl.User
	if tx, ok := getTX(ctx); ok {
		uc = tx.User
	}
	u, err := uc.Create().
		SetUUID(xid.New()).
		SetUserID(userID).
		SetPassword(password).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	user := &entity.User{
		ID:       u.UUID.String(),
		UserID:   u.UserID,
		Password: u.Password,
	}

	return user, nil
}
