package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/goji/context"
	"github.com/zenazn/goji/web"

	"github.com/andrew-d/flypaper/datastore"
)

func GetPort(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = context.FromC(c)
		idStr = c.URLParams["port"]
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	port, err := datastore.GetPort(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(port)
}
