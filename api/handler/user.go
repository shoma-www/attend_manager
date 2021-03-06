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
	ss     service.Session
}

// NewUser コンストラクタ
func NewUser(l core.Logger, user *service.User, session service.Session) *User {
	return &User{
		logger: l,
		us:     user,
		ss:     session,
	}
}

// Register ユーザ登録
func (u *User) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.logger.WithUUID(ctx).Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var uf UserForm
	if err := json.Unmarshal(req, &uf); err != nil {
		u.logger.WithUUID(ctx).Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := u.us.Register(ctx, uf.convert()); err != nil {
		u.logger.WithUUID(ctx).Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// SignIn サインイン
func (u *User) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := core.GetLogger(ctx)
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var form UserSignInForm
	if err := d.Decode(&form); err != nil {
		l.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ud, err := u.us.SigIn(ctx, form.GroupName, form.LoginID, form.Password)
	if err != nil {
		l.Error(err.Error())
		if err == entity.ErrUnauthenticated {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s := make(entity.Store)
	s[entity.GroupNameKey] = ud.GroupName
	s[entity.UserIDKey] = ud.UserID
	s[entity.UserNameKey] = ud.UserName

	ss, err := u.ss.Start(ctx, s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     service.SessionDataName,
		Value:    string(ss.ID),
		Domain:   "attend-manager.localhost",
		Secure:   true,
		HttpOnly: true,
	})
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

// UserSignInForm SignInフォーム
type UserSignInForm struct {
	GroupName string `json:"group_name"`
	LoginID   string `json:"login_id"`
	Password  string `json:"password"`
}
