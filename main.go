package main

import (
	"net/http"
	"time"

	webcontext "github.com/goji/context"
	"github.com/tylerb/graceful"
	"github.com/zenazn/goji/web"
	"golang.org/x/net/context"

	"github.com/andrew-d/flypaper/conf"
	"github.com/andrew-d/flypaper/middleware"
	"github.com/andrew-d/flypaper/router"
)

func main() {
	mux := router.New()

	mux.Use(middleware.Options)
	mux.Use(ContextMiddleware)
	mux.Use(middleware.SetHeaders)

	graceful.Run(conf.C.HostString(), 10*time.Second, mux)
}

func ContextMiddleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		// TODO: fill in context here

		// Add the context to the goji web context
		webcontext.Set(c, ctx)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
