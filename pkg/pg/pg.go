package pg

import (
	"context"
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
