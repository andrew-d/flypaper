package model

type Host struct {
	ID        int64  `db:"id"        json:"id"`
	IpAddress string `db:"ipaddress" json:"ipaddress"`
	Hostname  string `db:"hostname"  json:"hostname"`

	// TODO: marking a host as "standby", so it's still saved but not shown,
	// scanned by default, etc.?
}
