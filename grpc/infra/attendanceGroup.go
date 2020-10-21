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
	g, err := cl.Query().Where(attendancegroup.NameEQ(name)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, entity.ErrAttendanceGroupNotFound
		}
		return nil, err
	}

	return &entity.AttendanceGroup{
		ID:   g.ID,
		Name: *g.Name,
	}, nil
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
