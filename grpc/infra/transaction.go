package infra

import (
	"context"
	"database/sql"

	"github.com/shoma-www/attend_manager/grpc/ent"
)

var txKey = struct{}{}

type transaction struct {
	cl *ent.Client
}

func (t *transaction) Transaction(
	ctx context.Context, target func(tctx context.Context) error) error {
	tx, err := t.cl.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	txCtx := context.WithValue(ctx, txKey, tx)
	if err := target(txCtx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func getTX(ctx context.Context) (*ent.Tx, bool) {
	tx, ok := ctx.Value(txKey).(*ent.Tx)
	return tx, ok
}
