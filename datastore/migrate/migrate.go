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

func (m Migrator) Setup(tx migration.LimitedTx) error {
	stmts := []string{
		hostsTable,
		portsTable,
	}

	for _, stmt := range stmts {
		if _, err := tx.Exec(stmt); err != nil {
			return err
		}
	}

	return nil
}

const hostsTable = `
CREATE TABLE IF NOT EXISTS hosts (
	id        INTEGER PRIMARY KEY AUTOINCREMENT,
	ipaddress TEXT NOT NULL,
	hostname  TEXT
)
`
const portsTable = `
CREATE TABLE IF NOT EXISTS ports (
	id        INTEGER PRIMARY KEY AUTOINCREMENT,
	port      INTEGER NOT NULL
)
`
