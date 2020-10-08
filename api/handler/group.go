package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/shoma-www/attend_manager/api/entity"
	"github.com/shoma-www/attend_manager/api/service"
	"github.com/shoma-www/attend_manager/core"
)

// Group is handler
type Group struct {
	gs service.Group
}

// NewGroup コンストラクタ
func NewGroup(gs service.Group) *Group {
	return &Group{gs: gs}
}

// Create Group
func (g *Group) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := core.GetLogger(ctx)
	var gf GroupForm
	if err := json.NewDecoder(r.Body).Decode(&gf); err != nil {
		err = errors.Wrap(err, "invalid body paramaters")
		l.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	group, user := gf.convert()
	if err := g.gs.Create(ctx, group, user); err != nil {
		err = errors.Wrap(err, "failer create group")
		l.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GroupForm フォーム
type GroupForm struct {
	Name     string `json:"name"`
	LoginID  string `json:"login_id"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
}

func (gf GroupForm) convert() (entity.Group, entity.User) {
	g := entity.Group{
		Name: gf.Name,
	}
	u := entity.User{
		LoginID:  gf.LoginID,
		Password: gf.Password,
		Name:     gf.UserName,
	}
	return g, u
}
