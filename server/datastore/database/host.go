package database

import (
	"github.com/jmoiron/sqlx"

	"github.com/andrew-d/flypaper/server/model"
)

type HostStore struct {
	db *sqlx.DB
}

func NewHostStore(db *sqlx.DB) *HostStore {
	return &HostStore{db}
}

func (s *HostStore) GetHost(id int64) (*model.Host, error) {
	host := &model.Host{}
	err := s.db.Get(host, s.db.Rebind(hostGetQuery), id)
	return host, err
}

func (s *HostStore) PostHost(host *model.Host) error {
	ret, err := s.db.Exec(RebindInsert(s.db, hostInsertQuery),
		host.IpAddress,
		host.Hostname,
	)
	if err != nil {
		return err
	}

	host.ID, _ = ret.LastInsertId()
	return nil
}

const hostGetQuery = `
SELECT *
FROM hosts
WHERE id = ?
`

const hostInsertQuery = `
INSERT
INTO hosts (
	 ipaddress
	,hostname
)
VALUES (?, ?)
`
