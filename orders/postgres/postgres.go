package postgres

//go:generate sqlboiler --wipe psql --no-hooks --no-auto-timestamps

import (
	"context"
	"database/sql"
	"time"

	"github.com/ademuanthony/bitenvoy/orders/proto/orders"
	"github.com/ademuanthony/bitenvoy/utils"
	"github.com/micro/go-micro/util/log"
)

func NewPgDb(host, port, user, pass, dbname string) (*PgDb, error) {
	db, err := utils.PgConnect(host, port, user, pass, dbname)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(5)
	return &PgDb{
		db: db,
		queryTimeout: time.Second * 30,
	}, nil
}

type PgDb struct {
	db *sql.DB
	queryTimeout time.Duration
}

func (pg *PgDb) CreateOrder(ctx context.Context, order go_micro_srv_orders.Order) error {
	panic("implement me")
}

func (pg *PgDb) Close() error {
	log.Trace("Closing postgresql connection")
	return pg.db.Close()
}
