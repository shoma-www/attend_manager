package infra

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/xid"
	"github.com/shoma-www/attend_manager/ent"
	"github.com/shoma-www/attend_manager/ent/user"
	"github.com/shoma-www/attend_manager/grpc/entity"
)

type userDAO struct {
}

func (userDAO) Get(ctx context.Context, userID string) ([]*entity.User, error) {
	client, err := ent.Open("mysql", "root:root@tcp(mysql:3306)/attend")
	if err != nil {
		return nil, err
	}
	defer client.Close()
	us, _ := client.User.Query().Where(user.UserIDEQ(userID)).All(ctx)
	fmt.Printf("%v", us)
	users := make([]*entity.User, 0, len(us))
	for _, u := range us {
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

func (userDAO) Register(ctx context.Context, userID, password string) (*entity.User, error) {
	guid := xid.New()

	client, err := ent.Open("mysql", "root:root@tcp(mysql:3306)/attend")
	if err != nil {
		return nil, err
	}
	defer client.Close()
	u, err := client.Debug().User.Create().
		SetUUID(guid).
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
