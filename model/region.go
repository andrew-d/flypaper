package model

type Region struct {
	ID   int64  `db:"id"         json:"id"`
	Name string `db:"name"       json:"name"`

	// Times (in UTC) during which it is allowed to perform
	// tests on this region.
	TestStart NullTime `db:"test_start" json:"test_start"`
	TestEnd   NullTime `db:"test_end"   json:"test_end"`
}
