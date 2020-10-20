package infra

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/ent"
)

var txKey = struct{}{}

type transaction struct {
	l  core.Logger
	cl *ent.Client
}

func (t *transaction) Transaction(
	ctx context.Context, target func(tctx context.Context) (interface{}, error)) (interface{}, error) {
	t.l.Debug("start transaction\n")
	tx, err := t.cl.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "falied begin transaction")
	}
	txCtx := context.WithValue(ctx, txKey, tx)
	v, err := target(txCtx)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, errors.Wrap(err, "rollback failed")
		}
		t.l.Debug("rollback transaction\n")
		return nil, errors.Wrap(err, "rollback transaction")
	}
	t.l.Debug("finish transaction\n")
	return v, tx.Commit()
}

func getTX(ctx context.Context) (*ent.Tx, bool) {
	tx, ok := ctx.Value(txKey).(*ent.Tx)
	return tx, ok
}
