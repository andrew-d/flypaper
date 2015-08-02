package router

import (
	"github.com/zenazn/goji/web"

	"github.com/andrew-d/flypaper/server/handler"
)

func New() *web.Mux {
	mux := web.New()

	mux.Get("/api/regions/:region", handler.GetRegion)
	mux.Post("/api/regions", handler.PostRegion)

	mux.Get("/api/hosts/:host", handler.GetHost)
	mux.Post("/api/hosts", handler.PostHost)

	mux.Get("/api/ports/:port", handler.GetPort)

	return mux
}
