package model

import (
	"database/sql"
	"encoding/json"
)

type Region struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`

	// Times (in UTC) during which it is allowed to perform
	// tests on this region.
	TestStart sql.NullInt64 `db:"test_start"`
	TestEnd   sql.NullInt64 `db:"test_end"`
}

func (r *Region) MarshalJSON() ([]byte, error) {
	ret := struct {
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		TestStart *int64 `json:"test_end,omitempty"`
		TestEnd   *int64 `json:"test_start,omitempty"`
	}{
		ID:   r.ID,
		Name: r.Name,
	}

	if r.TestStart.Valid {
		ret.TestStart = &r.TestStart.Int64
	}
	if r.TestEnd.Valid {
		ret.TestEnd = &r.TestEnd.Int64
	}

	return json.Marshal(ret)
}
