package datastore

import (
	"golang.org/x/net/context"

	"github.com/andrew-d/flypaper/server/model"
)

type RegionStore interface {
	// GetRegion retrieves a region from the datastore for the given ID.
	GetRegion(id int64) (*model.Region, error)

	// PostRegion saves a new region in the datastore.
	PostRegion(region *model.Region) error
}

func GetRegion(c context.Context, id int64) (*model.Region, error) {
	return FromContext(c).GetRegion(id)
}

func PostRegion(c context.Context, region *model.Region) error {
	return FromContext(c).PostRegion(region)
}
