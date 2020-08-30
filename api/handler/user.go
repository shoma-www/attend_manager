package handler

import (
	"net/http"

	"github.com/shoma-www/attend_manager/api/infra"
	"github.com/shoma-www/attend_manager/core"
)

// User is handler
type User struct {
	logger  core.Logger
	factory *infra.RepoFactory
}

// NewUser コンストラクタ
func NewUser(l core.Logger, f *infra.RepoFactory) *User {
	return &User{logger: l, factory: f}
}

// Register ユーザ登録
func (u *User) Register(w http.ResponseWriter, r *http.Request) {
	u.logger.WithUUID(r.Context()).Debug("test")
	w.WriteHeader(http.StatusNotFound)
}
