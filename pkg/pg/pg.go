package pg

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgPool struct {
	*pgxpool.Pool
}

func NewPgPool(config *pgxpool.Config) (*PgPool, error) {
	dbPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to create new pool: %w", err)
	}
	return &PgPool{dbPool}, nil
}

func (db *PgPool) ConnString() string {
	return db.Config().ConnConfig.ConnString()
}

func (db *PgPool) ExecuteInTransaction(ctx context.Context, fn func(pgx.Tx) error) error {
	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				klog.CtxErrorf(ctx, "Failed to rollback transaction: %v", rbErr)
			}
		} else {
			if commitErr := tx.Commit(ctx); commitErr != nil {
				klog.CtxErrorf(ctx, "Failed to commit transaction: %v", commitErr)
				err = commitErr
			}
		}
	}()

	if err = fn(tx); err != nil {
		return fmt.Errorf("transaction execution failed: %w", err)
	}

	return nil
}
