package postgres

import (
	"fmt"

	"github.com/micro/go-micro/util/log"
)

const (
	createProviderTable = `CREATE TABLE IF NOT EXISTS provider (
		id VARCHAR(64) NOT NULL PRIMARY KEY,
		name VARCHAR(64) NOT NULL,
		logo VARCHAR(265) NOT NULL,
		status INT NOT NULL,
		discount FLOAT8 NOT NULL);`

	createCountryTable = `CREATE TABLE IF NOT EXISTS country (
		id VARCHAR(64) NOT NULL,
		name VARCHAR(64) NOT NULL,
		PRIMARY KEY(id)
	);`

	createProviderCountryTable = `CREATE TABLE IF NOT EXISTS provider_country (
		provider_id VARCHAR(64) NOT NULL,
		country_id VARCHAR(64) NOT NULL,
		PRIMARY KEY(provider_id, country_id)
	);`

	createHistoryTable = `CREATE TABLE IF NOT EXISTS history (
		id VARCHAR(64) NOT NULL PRIMARY KEY,
		username VARCHAR(64) NOT NULL,
		date INTEGER NOT NULL,
		phone_number VARCHAR(265) NOT NULL,
		network VARCHAR(265) NOT NULL,
		amount FLOAT8 NOT NULL);`

)

func (pg *PgDb) CreateProviderTable() error {
	log.Trace("Creating provider table")
	_, err := pg.db.Exec(createProviderTable)
	return err
}

func (pg *PgDb) ProviderTableExists() bool {
	exists, _ := pg.tableExists("provider")
	return exists
}

func (pg *PgDb) CreateCountryTable() error {
	log.Trace("Creating country table")
	_, err := pg.db.Exec(createCountryTable)
	return err
}

func (pg *PgDb) CountryTableExists() bool {
	exists, _ := pg.tableExists("country")
	return exists
}

func (pg *PgDb) CreateProviderCountryTable() error {
	log.Trace("Creating provider country table")
	_, err := pg.db.Exec(createProviderCountryTable)
	return err
}

func (pg *PgDb) ProviderCountryTableExists() bool {
	exists, _ := pg.tableExists("provider_country")
	return exists
}

func (pg *PgDb) CreateHistoryTable() error {
	log.Trace("Creating history country table")
	_, err := pg.db.Exec(createHistoryTable)
	return err
}

func (pg *PgDb) HistoryTableExists() bool {
	exists, _ := pg.tableExists("history")
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
	if err := pg.dropIndex("provider_country"); err != nil {
		return err
	}

	if err := pg.dropIndex("provider"); err != nil {
		return err
	}

	if err := pg.dropIndex("country"); err != nil {
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

