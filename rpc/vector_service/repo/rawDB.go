package repo

import (
	"cloud_tinamic/config"
	. "cloud_tinamic/pkg/errors"
	"cloud_tinamic/pkg/pg"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	db *pg.PgPool //定义一个连接池
)

func NewPgPool() *pg.PgPool {
	cfg := conf.GetConfigInstance()
	constr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.GetString("database.postgresql.user"),
		cfg.GetString("database.postgresql.password"),
		cfg.GetString("database.postgresql.host"),
		cfg.GetInt32("database.postgresql.port"),
		cfg.GetString("database.postgresql.database"),
		cfg.GetString("database.postgresql.sslmode"))
	dbConfig := pg.NewPgConfig(pg.WithConnString(constr))
	db, _ = pg.NewPgPool(dbConfig.Config)
	klog.Infof("success connect db @ %s", cfg.GetString("database.postgresql.host"))
	return db
}

func DBTileRequest(ctx context.Context, tr *TileRequest) ([]byte, error) {
	row := db.QueryRow(ctx, tr.SQL, tr.Args...)
	var mvtTile []byte
	err := row.Scan(&mvtTile)
	if err != nil {
		klog.Warn(err)
		// check for errors retrieving the rendered tile from the database
		// Timeout errors can occur if the context deadline is reached
		// or if the context is canceled during/before a database query.
		if pgconn.Timeout(err) {
			return nil, Kerrorf(TimeoutCode, "Timeout: deadline exceeded on %s/%s", tr.LayerID, tr.Tile.String())
		}
		return nil, Kerrorf(QueryFailedCode, "SQL error on %s/%s", tr.LayerID, tr.Tile.String())
	}
	return mvtTile, nil
}
