package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/goji/context"
	"github.com/zenazn/goji/web"

	"github.com/andrew-d/flypaper/server/datastore"
	"github.com/andrew-d/flypaper/server/log"
	"github.com/andrew-d/flypaper/server/model"
)

// ListRegions accepts a request to retrieve a list of regions.
//
//     GET /api/regions
//
func ListRegions(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = context.FromC(c)
		limit  = ToLimit(r)
		offset = ToOffset(r)
	)

	regions, err := datastore.ListRegions(ctx, limit, offset)
	if err != nil {
		log.FromContext(ctx).WithField("err", err).Error("Error listing regions")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(regions)
}

// GetRegion accepts a request to retrieve information about a particular region.
//
//     GET /api/regions/:region
//
func GetRegion(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = context.FromC(c)
		idStr = c.URLParams["region"]
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	region, err := datastore.GetRegion(ctx, id)
	if err != nil {
		log.FromContext(ctx).WithField("err", err).Error("Error getting region")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(region)
}

// PutRegion accepts a request to retrieve information about a particular region.
//
//     PUT /api/regions/:region
//
func PutRegion(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = context.FromC(c)
		idStr  = c.URLParams["region"]
		region model.Region
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !regionFromRequest(c, w, r, &region) {
		return
	}
	region.ID = id

	err = datastore.UpdateRegion(ctx, &region)
	if err != nil {
		log.FromContext(ctx).WithField("err", err).Error("Error updating region")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&region)
}

// DeleteRegion accepts a request to delete a region.
//
//     DELETE /api/regions/:region
//
func DeleteRegion(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = context.FromC(c)
		idStr = c.URLParams["region"]
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = datastore.DeleteRegion(ctx, id)
	if err != nil {
		log.FromContext(ctx).WithField("err", err).Error("Error deleting region")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// PostRegion accepts a request to add a new region.
//
//     POST /api/regions
//
func PostRegion(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = context.FromC(c)
		region model.Region
	)

	if !regionFromRequest(c, w, r, &region) {
		return
	}

	err := datastore.InsertRegion(ctx, &region)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&region)
}

// Attempts to deserialize a region from the incoming request.  If there is an
// error, writes a message to the response and returns false.
func regionFromRequest(c web.C, w http.ResponseWriter, r *http.Request, region *model.Region) bool {
	// Unmarshal the region from the payload
	defer r.Body.Close()
	in := struct {
		Name      string `json:"name"`
		TestStart *int64 `json:"test_start"`
		TestEnd   *int64 `json:"test_end"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}

	// Validate input
	if len(in.Name) < 1 {
		http.Error(w, "no name given", http.StatusBadRequest)
		return false
	}
	if (in.TestStart != nil) != (in.TestEnd != nil) {
		http.Error(w, "if a test start or end is given, both must be provided", http.StatusBadRequest)
		return false
	}

	// Create our 'normal' model.
	*region = model.Region{}
	region.Name = in.Name

	if in.TestStart != nil {
		region.TestStart.Valid = true
		region.TestStart.Int64 = *in.TestStart
	}
	if in.TestEnd != nil {
		region.TestEnd.Valid = true
		region.TestEnd.Int64 = *in.TestEnd
	}

	return true
}
