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

func NewPgPool(config *pgxpool.Config) *PgPool {
	dbPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		// TODO 错误处理
		return nil
	}
	return &PgPool{
		dbPool,
	}
}

func (db *PgPool) ConnString() string {
	return db.Config().ConnConfig.ConnString()
}

func (db *PgPool) ExecuteInTransaction(ctx context.Context, fn func(tx pgx.Tx) error) error {
	// 开启事务
	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		klog.Errorf("Failed to begin transaction: %v", err)
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// 确保事务在完成时要么提交要么回滚
	defer func() {
		if err != nil {
			// 回滚事务
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				klog.Errorf("Failed to rollback transaction: %v", rbErr)
			} else {
				klog.Infof("Transaction rolled back successfully")
			}
		} else {
			// 提交事务
			if commitErr := tx.Commit(ctx); commitErr != nil {
				klog.Errorf("Failed to commit transaction: %v", commitErr)
				err = commitErr // 捕获提交错误
			} else {
				klog.Infof("Transaction committed successfully")
			}
		}
	}()

	// 执行传入的事务逻辑
	err = fn(tx)
	if err != nil {
		klog.Errorf("Transaction execution failed: %v", err)
		return fmt.Errorf("transaction execution failed: %w", err)
	}

	return nil
}
