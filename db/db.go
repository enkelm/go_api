package db

import (
	"context"
	"log"
	"sync"

	"github.com/enkelm/go_api/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *postgres
	pgOnce     sync.Once
)

func InitMigration() {
	util.ExecuteCommand("tern", "migrate", "-m", "db/migrations", "-c", "db/migrations/tern.conf")
}

func NewPGInstance(ctx context.Context) (*postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, util.DB_CONN_STR)
		if err != nil {
			log.Println("Unable to connect to Postgres Db: %w", err)
			return
		}
		pgInstance = &postgres{db}
	})

	return pgInstance, nil
}

func (pg *postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *postgres) Close() {
	pg.db.Close()
}
