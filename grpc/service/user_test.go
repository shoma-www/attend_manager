package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/entity"
	"github.com/shoma-www/attend_manager/grpc/mock_service"
)

type mockTransaction struct{}

func (mockTransaction) Transaction(ctx context.Context, target func(tctx context.Context) error) error {
	return target(ctx)
}

func TestUser_Register(t *testing.T) {
	l := core.NewLogger(core.Debug)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("success", func(t *testing.T) {
		u := &entity.User{
			UserID:   "user success",
			Password: "password",
		}

		ctx := context.Background()
		mr := mock_service.NewMockUserRepository(ctrl)
		mr.EXPECT().Get(ctx, u.UserID).Return(nil, entity.ErrUserNotFound)
		mr.EXPECT().Register(ctx, u.UserID, gomock.Any()).Return(u, nil)

		us := NewUser(l, mockTransaction{}, mr)
		if err := us.Register(ctx, u.UserID, u.Password); err != nil {
			t.Errorf("User.Register() error = %v, wantErr nil", err)
		}
	})

	t.Run("duplicate error", func(t *testing.T) {
		u := &entity.User{
			UserID:   "user error",
			Password: "password",
		}

		ctx := context.Background()
		wamtErr := entity.ErrDuplicatedUser
		mr := mock_service.NewMockUserRepository(ctrl)
		mr.EXPECT().Get(ctx, u.UserID).Return([]*entity.User{u}, nil)

		us := NewUser(l, mockTransaction{}, mr)
		if err := us.Register(ctx, u.UserID, u.Password); err != wamtErr {
			t.Errorf("User.Register() error = %v, wantErr %v", err, wamtErr)
		}
	})

	t.Run("other error", func(t *testing.T) {
		u := &entity.User{
			UserID:   "user other error",
			Password: "password",
		}

		ctx := context.Background()
		wamtErr := errors.New("other error")
		mr := mock_service.NewMockUserRepository(ctrl)
		mr.EXPECT().Get(ctx, u.UserID).Return(nil, wamtErr)

		us := NewUser(l, mockTransaction{}, mr)
		if err := us.Register(ctx, u.UserID, u.Password); err != wamtErr {
			t.Errorf("User.Register() error = %v, wantErr %v", err, wamtErr)
		}
	})
}
