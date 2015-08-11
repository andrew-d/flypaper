package database

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/andrew-d/flypaper/server/model"
)

type RegionStoreSuite struct {
	suite.Suite

	db *sqlx.DB
	rs *RegionStore
}

func TestRegionStore(t *testing.T) {
	db := MustConnect("sqlite3", ":memory:")
	rs := NewRegionStore(db)

	// Don't want to run this as part of our test suite, since we want to
	// test before we muck with the DB.
	regions, err := rs.ListRegions(10, 0)
	if assert.NoError(t, err) && assert.Equal(t, 1, len(regions)) {
		assert.Equal(t, "default", regions[0].Name)
	}

	suite.Run(t, &RegionStoreSuite{
		db: db,
		rs: rs,
	})
}

func (s *RegionStoreSuite) SetupTest() {
	s.db.Exec(`DELETE FROM regions`)
}

func (s *RegionStoreSuite) TestInsertRegion() {
	r := model.Region{
		Name: "foo",
	}
	err := s.rs.InsertRegion(&r)

	s.NoError(err)
	s.NotEqual(0, r.ID)
}

func (s *RegionStoreSuite) TestDeleteRegion() {
	r := model.Region{
		Name: "foo",
	}
	err := s.rs.InsertRegion(&r)

	regions, err := s.rs.ListRegions(10, 0)
	s.NoError(err)
	s.Equal(1, len(regions))

	err = s.rs.DeleteRegion(r.ID)
	s.NoError(err)

	regions, err = s.rs.ListRegions(10, 0)
	s.NoError(err)
	s.Equal(0, len(regions))
}

func (s *RegionStoreSuite) TestUpdateRegion() {
	r1 := model.Region{
		Name: "foo",
	}
	err := s.rs.InsertRegion(&r1)

	region, err := s.rs.GetRegion(r1.ID)
	s.NoError(err)
	s.Equal("foo", region.Name)

	err = s.rs.UpdateRegion(&model.Region{
		ID:   r1.ID,
		Name: "bar",
	})
	s.NoError(err)

	region, err = s.rs.GetRegion(r1.ID)
	s.NoError(err)
	s.Equal("bar", region.Name)
}
