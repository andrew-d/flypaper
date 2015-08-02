package database

import (
	"github.com/jmoiron/sqlx"

	"github.com/andrew-d/flypaper/server/model"
)

type RegionStore struct {
	db *sqlx.DB
}

func NewRegionStore(db *sqlx.DB) *RegionStore {
	return &RegionStore{db}
}

func (s *RegionStore) GetRegion(id int64) (*model.Region, error) {
	region := &model.Region{}
	err := s.db.Get(region, s.db.Rebind(regionGetQuery), id)
	return region, err
}

func (s *RegionStore) PostRegion(region *model.Region) error {
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
