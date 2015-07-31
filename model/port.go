package model

type Port struct {
	ID   int64  `db:"id"   json:"id"`
	Port uint16 `db:"port" json:"port"`
	Host int64  `db:"host" json:"-"`
}
