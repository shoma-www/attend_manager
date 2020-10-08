package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/shoma-www/attend_manager/api/entity"
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

	if err := u.us.Register(ctx, uf.convert()); err != nil {
		u.logger.WithUUID(ctx).Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// UserForm フォーム
type UserForm struct {
	GroupID  string `json:"group_id"`
	LoginID  string `json:"login_id"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (u UserForm) convert() entity.User {
	return entity.User{
		GroupID:  u.GroupID,
		LoginID:  u.LoginID,
		Password: u.Password,
		Name:     u.Name,
	}
}
