package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shoma-www/attend_manager/api/entity"
	"github.com/shoma-www/attend_manager/core"
)

// Group operation
type Group struct {
	gr GroupRepository
}

// NewGroup コンストラクタ
func NewGroup(gr GroupRepository) Group {
	return Group{gr: gr}
}

// Create グループの作成
func (g *Group) Create(ctx context.Context, group entity.Group, user entity.User) error {
	l := core.GetLogger(ctx)
	l.Info("Create Group: %s", group.Name)
	if err := g.gr.Create(ctx, group, user); err != nil {
		l.Error("Failed Create Group")
		return errors.Wrap(err, "Failed Create Group")
	}
	l.Info("Success Create Group")
	return nil
}
