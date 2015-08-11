package database

import (
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/andrew-d/flypaper/server/model"
)

type RegionStore struct {
	db *sqlx.DB
}

func NewRegionStore(db *sqlx.DB) *RegionStore {
	return &RegionStore{db}
}

func (s *RegionStore) ListRegions(limit, offset int) ([]*model.Region, error) {
	regions := []*model.Region{}
	err := s.db.Select(&regions, s.db.Rebind(regionListQuery), limit, offset)
	return regions, err
}

func (s *RegionStore) GetRegion(id int64) (*model.Region, error) {
	region := &model.Region{}
	err := s.db.Get(region, s.db.Rebind(regionGetQuery), id)
	return region, err
}

func (s *RegionStore) InsertRegion(region *model.Region) error {
	ret, err := s.db.Exec(RebindInsert(s.db, regionInsertQuery),
		region.Name,
		region.TestStart,
		region.TestEnd,
	)
	if err != nil {
		return err
	}

	region.ID, _ = ret.LastInsertId()
	return nil
}

func (s *RegionStore) DeleteRegion(id int64) error {
	// Get the region first, to see if it's the 'default' region
	region, err := s.GetRegion(id)
	if err != nil {
		return err
	}

	if region.Name == "default" {
		return errors.New("cannot remove the default region")
	}

	_, err = s.db.Exec(s.db.Rebind(regionDeleteQuery), region.ID)
	return err
}

func (s *RegionStore) UpdateRegion(region *model.Region) error {
	_, err := s.db.Exec(s.db.Rebind(regionUpdateQuery),
		region.Name,
		region.TestStart,
		region.TestEnd,
		region.ID,
	)
	return err
}

const regionListQuery = `
SELECT *
FROM regions
ORDER BY id DESC
LIMIT ? OFFSET ?
`

const regionGetQuery = `
SELECT *
FROM regions
WHERE id = ?
`

const regionInsertQuery = `
INSERT
INTO regions (
     name
    ,test_start
    ,test_end
)
VALUES (?, ?, ?)
`

const regionDeleteQuery = `
DELETE
FROM regions
WHERE id = ?
`

const regionUpdateQuery = `
UPDATE regions
SET
 name = ?
,test_start = ?
,test_end = ?
WHERE id = ?
`
