package datastore

import (
	"golang.org/x/net/context"

	"github.com/andrew-d/flypaper/server/model"
)

type RegionStore interface {
	// ListRegions retrieves all regions from the database, possibly with an
	// offset or limit provided.
	ListRegions(limit, offset int) ([]*model.Region, error)

	// GetRegion retrieves a region from the datastore for the given ID.
	GetRegion(id int64) (*model.Region, error)

	// PostRegion saves a new region in the datastore.
	PostRegion(region *model.Region) error

	// DeleteRegion removes a region from the datastore.
	DeleteRegion(id int64) error

	// UpdateRegion updates a region in the datastore.
	UpdateRegion(region *model.Region) error
}

func ListRegions(c context.Context, limit, offset int) ([]*model.Region, error) {
	return FromContext(c).ListRegions(limit, offset)
}

func GetRegion(c context.Context, id int64) (*model.Region, error) {
	return FromContext(c).GetRegion(id)
}

func PostRegion(c context.Context, region *model.Region) error {
	return FromContext(c).PostRegion(region)
}

func DeleteRegion(c context.Context, id int64) error {
	return FromContext(c).DeleteRegion(id)
}
