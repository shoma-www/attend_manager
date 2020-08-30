package handler

import (
	"net/http"

	"github.com/shoma-www/attend_manager/api/service"
	"github.com/shoma-www/attend_manager/core"
)

// User is handler
type User struct {
	logger core.Logger
	us     *service.User
}

// NewUser コンストラクタ
func NewUser(l core.Logger, user *service.User) *User {
	return &User{logger: l, us: user}
}

// Register ユーザ登録
func (u *User) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	u.logger.WithUUID(ctx).Debug("test")
	err := u.us.Register(ctx, "hoge", "hoge")
	u.logger.WithUUID(ctx).Error(err.Error())
	w.WriteHeader(http.StatusNotFound)
}
