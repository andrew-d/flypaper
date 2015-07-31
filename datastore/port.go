package datastore

import (
	"golang.org/x/net/context"

	"github.com/andrew-d/flypaper/model"
)

type PortStore interface {
	// GetPort retrieves a port from the datastore for the given ID.
	GetPort(id int64) (*model.Port, error)

	// GetPortsForHost retrieves all ports for the given host.
	GetPortsForHost(host *model.Host) ([]*model.Port, error)

	// GetPortsByNum retrieves all ports with the given port number.
	GetPortsByNum(num uint16) ([]*model.Port, error)
}

func GetPort(c context.Context, id int64) (*model.Port, error) {
	return FromContext(c).GetPort(id)
}

func GetPortsForHost(c context.Context, host *model.Host) ([]*model.Port, error) {
	return FromContext(c).GetPortsForHost(host)
}

func GetPortsByNum(c context.Context, num uint16) ([]*model.Port, error) {
	return FromContext(c).GetPortsByNum(num)
}
