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

// GetHost accepts a request to retrieve information about a particular host.
//
//     GET /api/hosts/:host
//
func GetHost(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = context.FromC(c)
		idStr = c.URLParams["host"]
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	host, err := datastore.GetHost(ctx, id)
	if err != nil {
		log.FromContext(ctx).WithField("err", err).Error("Error getting host")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(host)
}

// PostHost accepts a request to add a new host.
//
//     POST /api/hosts
//
func PostHost(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		ctx = context.FromC(c)
	)

	// Unmarshal the host from the payload
	defer r.Body.Close()
	in := struct {
		IpAddress string  `json:"ipaddress"`
		Hostname  *string `json:"hostname"`
		Region    *int64  `json:"region"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create our 'normal' model.
	host := new(model.Host)

	if len(in.IpAddress) < 1 {
		http.Error(w, "no IP address given", http.StatusBadRequest)
		return
	}

	host.IpAddress = in.IpAddress

	if in.Hostname != nil {
		host.Hostname.Valid = true
		host.Hostname.String = *in.Hostname
	}

	if in.Region != nil {
		host.Region.Valid = true
		host.Region.Int64 = *in.Region
	}

	err := datastore.PostHost(ctx, host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(host)
}
