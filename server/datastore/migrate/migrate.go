package migrate

import (
	"github.com/BurntSushi/migration"
	"github.com/jmoiron/sqlx"
)

type Migrator struct {
	DbType string
}

func (m Migrator) rebind(s string) string {
	return sqlx.Rebind(sqlx.BindType(m.DbType), s)
}

// Setup will create all necessary tables and indexes in the database.
func (m Migrator) Setup(tx migration.LimitedTx) error {
	stmts := []string{
		regionsTable,
		hostsTable,
		portsTable,
		portsTableIndex,
	}

	for _, stmt := range stmts {
		if _, err := tx.Exec(stmt); err != nil {
			return err
		}
	}

	return nil
}

// DefaultRegion will insert a default region into the database with the
// name "default", no start, and no end time.
func (m Migrator) DefaultRegion(tx migration.LimitedTx) error {
	_, err := tx.Exec(
		`INSERT INTO regions(name) VALUES(?)`,
		"default",
	)
	return err
}

const regionsTable = `
CREATE TABLE IF NOT EXISTS regions (
	 id         INTEGER PRIMARY KEY AUTOINCREMENT
	,name       TEXT NOT NULL
	,test_start INTEGER
	,test_end   INTEGER

	,UNIQUE(name)
)
`

const hostsTable = `
CREATE TABLE IF NOT EXISTS hosts (
	 id        INTEGER PRIMARY KEY AUTOINCREMENT
	,ipaddress TEXT NOT NULL
	,hostname  TEXT
	,region    INTEGER

	,FOREIGN KEY (region) REFERENCES regions(id)
	,UNIQUE(ipaddress)
)
`
const portsTable = `
CREATE TABLE IF NOT EXISTS ports (
	 id        INTEGER PRIMARY KEY AUTOINCREMENT
	,port      INTEGER NOT NULL
	,host      INTEGER NOT NULL

	,FOREIGN KEY (host) REFERENCES hosts(id)
	,UNIQUE(host, port)
)
`

const portsTableIndex = `
CREATE INDEX IF NOT EXISTS port_idx ON ports(port)
`
