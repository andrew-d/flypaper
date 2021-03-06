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

	"github.com/andrew-d/flypaper/server/conf"
	"github.com/andrew-d/flypaper/server/datastore"
	"github.com/andrew-d/flypaper/server/datastore/database"
	"github.com/andrew-d/flypaper/server/log"
	"github.com/andrew-d/flypaper/server/middleware"
	"github.com/andrew-d/flypaper/server/router"
)

// Generic structure that holds all created things - database connection,
// datastore, etc.
type Vars struct {
	db  *sqlx.DB
	ds  datastore.Datastore
	log *logrus.Logger
}

var (
	// Commit SHA and version for the current build, set by the
	// compile process.
	version  string
	revision string
)

func main() {
	var vars Vars

	// Create logger.
	vars.log = log.NewLogger()
	vars.log.Info("initializing...")

	// Connect to the database.
	db, err := database.Connect(conf.C.DbType, conf.C.DbConn)
	if err != nil {
		vars.log.WithFields(logrus.Fields{
			"err":     err,
			"db_type": conf.C.DbType,
			"db_conn": conf.C.DbConn,
		}).Error("Could not connect to database")
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

// ContextMiddleware will add our variables to the per-request context.
func ContextMiddleware(vars *Vars) web.MiddlewareType {
	mfn := func(c *web.C, h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.Background()
			ctx = datastore.NewContext(ctx, vars.ds)
			ctx = log.NewContext(ctx, vars.log)

			// Add the context to the goji web context
			webcontext.Set(c, ctx)
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}

	return mfn
}
