package entrepository

import (
	"context"

	"apps/investimento/pkg/ent"

	"apps/investimento/pkg/support/errors"
)

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return errors.Wrap("ent:support:withtx:client", err)
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			e := errors.Wrap("ent:support:withtx:rollback", rerr)
			err = errors.Wrap(e.Error(), err)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap("ent:support:withtx:commit", err)
	}
	return nil
}
