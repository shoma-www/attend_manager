package infra

import (
	"context"

	"github.com/pkg/errors"
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

func (ud *userDAO) Get(ctx context.Context, groupID xid.ID, loginID string) (*entity.User, error) {
	gr := ud.cl.AttendanceGroup
	if tx, ok := getTX(ctx); ok {
		gr = tx.AttendanceGroup
	}
	group, err := gr.Get(ctx, groupID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get attendance groups in users")
	}

	u, err := group.QueryUsers().Where(user.LoginID(loginID)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, entity.ErrUserNotFound
		}
		return nil, errors.Wrap(err, "failed to get users")
	}
	ud.logger.Debug("id: %s, group ID:%s, loginID: %s, pass: %s", u.ID, groupID, u.LoginID, u.Password)
	user1 := &entity.User{
		ID:       u.ID,
		GroupID:  groupID,
		LoginID:  u.LoginID,
		Password: u.Password,
	}
	if u.Name != nil {
		user1.Name = *u.Name
	}
	return user1, nil
}

func (ud *userDAO) Register(ctx context.Context, groupID xid.ID, logiinID, password, name string) (*entity.User, error) {
	uc := ud.cl.User
	if tx, ok := getTX(ctx); ok {
		uc = tx.User
	}
	u, err := uc.Create().
		SetID(xid.New()).
		SetGroupID(groupID).
		SetLoginID(logiinID).
		SetPassword(password).
		SetName(name).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:       u.ID,
		GroupID:  groupID,
		LoginID:  u.LoginID,
		Name:     name,
		Password: u.Password,
	}, nil
}
