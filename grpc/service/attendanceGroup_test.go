package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/entity"
)

func TestAttendanceGroup_Create(t *testing.T) {
	type fields struct {
		logger core.Logger
		tr     Transaction
		gr     AttendanceGroupRepository
		ur     UserRepository
	}
	type args struct {
		ctx       context.Context
		groupName string
		loginID   string
		password  string
		userName  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.AttendanceGroup
		want1   *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &AttendanceGroup{
				logger: tt.fields.logger,
				tr:     tt.fields.tr,
				gr:     tt.fields.gr,
				ur:     tt.fields.ur,
			}
			got, got1, err := g.Create(tt.args.ctx, tt.args.groupName, tt.args.loginID, tt.args.password, tt.args.userName)
			if (err != nil) != tt.wantErr {
				t.Errorf("AttendanceGroup.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AttendanceGroup.Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AttendanceGroup.Create() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
