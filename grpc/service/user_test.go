package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/rs/xid"
	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/entity"
	"github.com/shoma-www/attend_manager/grpc/mock_service"
)

type mockTransaction struct{}

func (mockTransaction) Transaction(ctx context.Context, target func(tctx context.Context) (interface{}, error)) (interface{}, error) {
	return target(ctx)
}

func TestUser_Register(t *testing.T) {
	l := core.NewLogger(core.Debug)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("success", func(t *testing.T) {
		u := &entity.User{
			GroupID:  xid.New(),
			LoginID:  "user success",
			Password: "password",
			Name:     "user name",
		}

		ctx := context.Background()
		mr := mock_service.NewMockUserRepository(ctrl)
		mr.EXPECT().Get(ctx, u.GroupID, u.LoginID).Return(nil, entity.ErrUserNotFound)
		mr.EXPECT().Register(ctx, u.GroupID, u.LoginID, gomock.Any(), u.Name).DoAndReturn(
			func(ctx context.Context, groupID xid.ID, loginID, password, name string) (*entity.User, error) {
				return &entity.User{
					GroupID:  groupID,
					LoginID:  loginID,
					Password: password,
					Name:     name,
				}, nil
			})

		us := NewUser(l, mockTransaction{}, mr)
		user, err := us.Register(ctx, u.GroupID, u.LoginID, u.Password, u.Name)
		if err != nil {
			t.Errorf("User.Register() error = %v, wantErr nil", err)
		}
		if diff := cmp.Diff(user, u, cmpopts.IgnoreFields(entity.User{}, "Password")); diff != "" {
			t.Errorf("Register differs:\n%s", diff)
		}
	})

	t.Run("duplicate error", func(t *testing.T) {
		u := &entity.User{
			GroupID:  xid.New(),
			LoginID:  "user error",
			Password: "password",
			Name:     "user error",
		}

		ctx := context.Background()
		wantErr := entity.ErrDuplicatedUser
		mr := mock_service.NewMockUserRepository(ctrl)
		mr.EXPECT().Get(ctx, u.GroupID, u.LoginID).Return(u, nil)

		us := NewUser(l, mockTransaction{}, mr)
		if _, err := us.Register(ctx, u.GroupID, u.LoginID, u.Password, u.Name); err != wantErr {
			t.Errorf("User.Register() error = %v, wantErr %v", err, wantErr)
		}
	})

	t.Run("other error", func(t *testing.T) {
		u := &entity.User{
			GroupID:  xid.New(),
			LoginID:  "user other error",
			Password: "password",
			Name:     "user other error",
		}

		ctx := context.Background()
		wantErr := errors.New("other error")
		mr := mock_service.NewMockUserRepository(ctrl)
		mr.EXPECT().Get(ctx, u.GroupID, u.LoginID).Return(nil, wantErr)

		us := NewUser(l, mockTransaction{}, mr)
		if _, err := us.Register(ctx, u.GroupID, u.LoginID, u.Password, u.Name); err != wantErr {
			t.Errorf("User.Register() error = %v, wantErr %v", err, wantErr)
		}
	})
}
