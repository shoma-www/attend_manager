package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/shoma-www/attend_manager/api/service"
	"github.com/shoma-www/attend_manager/core"
)

// UserForm フォーム
type UserForm struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

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
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.logger.WithUUID(ctx).Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var uf UserForm
	if err := json.Unmarshal(req, &uf); err != nil {
		u.logger.WithUUID(ctx).Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := u.us.Register(ctx, uf.UserID, uf.Password); err != nil {
		u.logger.WithUUID(ctx).Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusNotFound)
	}
}
