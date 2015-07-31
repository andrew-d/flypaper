package database

import (
	"github.com/jmoiron/sqlx"

	"github.com/andrew-d/flypaper/server/model"
)

type PortStore struct {
	db *sqlx.DB
}

func NewPortStore(db *sqlx.DB) *PortStore {
	return &PortStore{db}
}

func (s *PortStore) GetPort(id int64) (*model.Port, error) {
	port := &model.Port{}
	err := s.db.Get(&port, s.db.Rebind(portGetQuery), id)
	return port, err
}

func (s *PortStore) GetPortsForHost(host *model.Host) ([]*model.Port, error) {
	var ports []*model.Port
	err := s.db.Select(&ports, s.db.Rebind(portGetForHostQuery), host.ID)
	return ports, err
}

func (s *PortStore) GetPortsByNum(num uint16) ([]*model.Port, error) {
	var ports []*model.Port
	err := s.db.Select(&ports, s.db.Rebind(portGetByNumQuery), num)
	return ports, err
}

const portGetQuery = `
SELECT *
FROM ports
WHERE id = ?
`

const portGetForHostQuery = `
SELECT *
FROM ports
WHERE host = ?
`

const portGetByNumQuery = `
SELECT *
FROM ports
WHERE port = ?
`
