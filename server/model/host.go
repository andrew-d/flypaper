package model

import (
	"database/sql"
)

type Host struct {
	ID        int64          `db:"id"        json:"id"`
	IpAddress string         `db:"ipaddress" json:"ipaddress"`
	Hostname  sql.NullString `db:"hostname"  json:"hostname"`
	Region    sql.NullInt64  `db:"region"    json:"region"`

	// TODO: marking a host as "standby", so it's still saved but not shown,
	// scanned by default, etc.?
}
