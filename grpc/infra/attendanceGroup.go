package infra

import (
	"context"

	"github.com/rs/xid"
	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/ent"
	"github.com/shoma-www/attend_manager/grpc/ent/attendancegroup"
	"github.com/shoma-www/attend_manager/grpc/entity"
)

type attendanceGroupDAO struct {
	logger core.Logger
	cl     *ent.Client
}

func (ag *attendanceGroupDAO) Get(ctx context.Context, name string) (*entity.AttendanceGroup, error) {
	cl := ag.cl.AttendanceGroup
	groups, err := cl.Query().Where(attendancegroup.NameEQ(name)).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(groups) == 0 {
		return nil, entity.ErrAttendanceGroupNotFound
	}
	group := new(entity.AttendanceGroup)
	if groups[0].Name != nil {
		group.Name = *groups[0].Name
	}
	return group, nil
}

func (ag *attendanceGroupDAO) Create(ctx context.Context, name string) (*entity.AttendanceGroup, error) {
	cl := ag.cl.AttendanceGroup
	if tx, ok := getTX(ctx); ok {
		cl = tx.AttendanceGroup
	}
	g, err := cl.Create().
		SetID(xid.New()).
		SetName(name).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	group := &entity.AttendanceGroup{
		ID:   g.ID,
		Name: name,
	}

	return group, nil
}
