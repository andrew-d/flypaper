package main

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/andrew-d/webhelpers"
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
	db  *sqlx.DB
	ds  datastore.Datastore
	log *logrus.Logger
}

func main() {
	var vars Vars

	// Create logger.
	// TODO: depending on conf.C.Debug, we can set this to print JSON, etc.
	vars.log = logrus.New()
	vars.log.Info("initializing...")

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

	mux.Use(webhelpers.Recoverer)
	mux.Use(middleware.Options)
	mux.Use(ContextMiddleware(&vars))
	mux.Use(middleware.SetHeaders)

	// We wrap the Request ID middleware and our logger 'outside' the mux, so
	// all requests (including ones that aren't matched by the router) get
	// logged.
	var handler http.Handler = mux
	handler = webhelpers.LogrusLogger(vars.log, handler)
	handler = webhelpers.RequestID(handler)

	// Start serving
	vars.log.Infof("starting server on: %s", conf.C.HostString())
	graceful.Run(conf.C.HostString(), 10*time.Second, handler)
	vars.log.Info("server finished")
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
