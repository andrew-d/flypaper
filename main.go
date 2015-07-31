package main

import (
	"net/http"
	"time"

	webcontext "github.com/goji/context"
	"github.com/jmoiron/sqlx"
	"github.com/tylerb/graceful"
	"github.com/zenazn/goji/web"
	"golang.org/x/net/context"

	"github.com/andrew-d/flypaper/conf"
	"github.com/andrew-d/flypaper/datastore"
	"github.com/andrew-d/flypaper/datastore/database"
	"github.com/andrew-d/flypaper/middleware"
	"github.com/andrew-d/flypaper/router"
)

// Generic structure that holds all created things - database connection,
// datastore, etc.
type Vars struct {
	db *sqlx.DB
	ds datastore.Datastore
}

func main() {
	var vars Vars

	// Connect to the database.
	db, err := database.Connect(conf.C.DbType, conf.C.DbConn)
	if err != nil {
		// TODO: error message
		return
	}
	vars.db = db

	// Create datastore.
	vars.ds = database.NewDatastore(db)

	// Create router and add middleware.
	mux := router.New()

	mux.Use(middleware.Options)
	mux.Use(ContextMiddleware(&vars))
	mux.Use(middleware.SetHeaders)

	graceful.Run(conf.C.HostString(), 10*time.Second, mux)
}

func ContextMiddleware(vars *Vars) web.MiddlewareType {
	mfn := func(c *web.C, h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.Background()
			ctx = datastore.NewContext(ctx, vars.ds)

			// Add the context to the goji web context
			webcontext.Set(c, ctx)
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}

	return mfn
}
