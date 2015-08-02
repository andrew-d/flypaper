package datastore

import (
	"golang.org/x/net/context"

	"github.com/andrew-d/flypaper/server/model"
)

type HostStore interface {
	// GetHost retrieves a host from the datastore for the given ID.
	GetHost(id int64) (*model.Host, error)

	// PostHost saves a new host in the datastore.
	PostHost(host *model.Host) error
}

func GetHost(c context.Context, id int64) (*model.Host, error) {
	return FromContext(c).GetHost(id)
}

func PostHost(c context.Context, host *model.Host) error {
	return FromContext(c).PostHost(host)
}
