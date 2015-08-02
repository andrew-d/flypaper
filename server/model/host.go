package model

import (
	"database/sql"
	"encoding/json"
)

type Host struct {
	ID        int64          `db:"id"`
	IpAddress string         `db:"ipaddress"`
	Hostname  sql.NullString `db:"hostname"`
	Region    sql.NullInt64  `db:"region"`

	// TODO: marking a host as "standby", so it's still saved but not shown,
	// scanned by default, etc.?
}

func (h *Host) MarshalJSON() ([]byte, error) {
	ret := struct {
		ID        int64   `json:"id"`
		IpAddress string  `json:"ipaddress"`
		Hostname  *string `json:"hostname,omitempty"`
		Region    *int64  `json:"region,omitempty"`
	}{
		ID:        h.ID,
		IpAddress: h.IpAddress,
	}

	if h.Hostname.Valid {
		ret.Hostname = &h.Hostname.String
	}
	if h.Region.Valid {
		ret.Region = &h.Region.Int64
	}

	return json.Marshal(ret)
}
