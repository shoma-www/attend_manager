package service

import (
	"context"

	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/entity"
)

// AttendanceGroup user service
type AttendanceGroup struct {
	logger core.Logger
	tr     Transaction
	gr     AttendanceGroupRepository
	ur     UserRepository
}

// NewAttendanceGroup service constructor
func NewAttendanceGroup(l core.Logger, tr Transaction, gr AttendanceGroupRepository, ur UserRepository) *AttendanceGroup {
	return &AttendanceGroup{
		logger: l,
		tr:     tr,
		gr:     gr,
		ur:     ur,
	}
}

// Create new group
func (g *AttendanceGroup) Create(ctx context.Context, groupName, loginID, password, userName string) (*entity.AttendanceGroup, *entity.User, error) {
	type result struct {
		group *entity.AttendanceGroup
		user  *entity.User
	}
	g.logger.WithUUID(ctx).Info("create new group. name: %s", groupName)
	v, err := g.tr.Transaction(ctx, func(tctx context.Context) (interface{}, error) {
		if group, err := g.gr.Get(tctx, groupName); err != entity.ErrAttendanceGroupNotFound {
			if group != nil {
				return nil, entity.ErrDuplicatedAttendanceGroup
			}
			return nil, err
		}

		group, err := g.gr.Create(tctx, groupName)
		if err != nil {
			return nil, err
		}

		hashedPassword, err := core.GenerateHashedPassword(password)
		if err != nil {
			return nil, err
		}
		user, err := g.ur.Register(tctx, group.ID, loginID, string(hashedPassword), userName)
		if err != nil {
			return nil, err
		}
		g.logger.WithUUID(tctx).Info("complete created group. group id: %s, user id: %s", group.ID, user.ID)
		return &result{
			group: group,
			user:  user,
		}, nil
	})

	if r, ok := v.(result); ok {
		return r.group, r.user, err
	}
	return nil, nil, err
}
