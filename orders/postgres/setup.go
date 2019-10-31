package postgres

import (
	"fmt"

	"github.com/micro/go-micro/util/log"
)

const (
	createOrderTable = `CREATE TABLE IF NOT EXISTS order_model (
		id VARCHAR(64) NOT NULL PRIMARY KEY,
		username VARCHAR(64) NOT NULL,
		product VARCHAR(265) NOT NULL,
		status VARCHAR(265) NOT NULL,
		details TEXT NOT NULL,
		date INT NOT NULL,
		amount FLOAT8 NOT NULL);`

	createOrderStatusTable = `CREATE TABLE IF NOT EXISTS order_status (
		id VARCHAR(64) NOT NULL PRIMARY KEY,
		order_id VARCHAR(64) NOT NULL,
		status VARCHAR(265) NOT NULL,
		date INT NOT NULL
		)`
)

func (pg *PgDb) CreateOrderTable() error {
	log.Trace("Creating order table")
	_, err := pg.db.Exec(createOrderTable)
	return err
}

func (pg *PgDb) OrderTableExists() bool {
	exists, _ := pg.tableExists("order")
	return exists
}

func (pg *PgDb) CreateOrderHistoryTable() error {
	log.Trace("Creating order table")
	_, err := pg.db.Exec(createOrderStatusTable)
	return err
}

func (pg *PgDb) OrderHistoryTableExists() bool {
	exists, _ := pg.tableExists("order_status")
	return exists
}

func (pg *PgDb) tableExists(name string) (bool, error) {
	rows, err := pg.db.Query(`SELECT relname FROM pg_class WHERE relname = $1`, name)
	if err == nil {
		defer func() {
			if e := rows.Close(); e != nil {
				log.Error("Close of Query failed: ", e)
			}
		}()
		return rows.Next(), nil
	}
	return false, err
}

func (pg *PgDb) DropAllTables() error {
	if err := pg.dropIndex("order_status"); err != nil {
		return err
	}

	if err := pg.dropIndex("order"); err != nil {
		return err
	}

	return nil
}

func (pg *PgDb) dropTable(name string) error {
	log.Tracef("Dropping table %s", name)
	_, err := pg.db.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS %s;`, name))
	return err
}

func (pg *PgDb) dropIndex(name string) error {
	log.Tracef("Dropping table %s", name)
	_, err := pg.db.Exec(fmt.Sprintf(`DROP INDEX IF EXISTS %s;`, name))
	return err
}

