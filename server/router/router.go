package router

import (
	"github.com/zenazn/goji/web"

	"github.com/andrew-d/flypaper/server/handler"
)

func New() *web.Mux {
	mux := web.New()

	mux.Get("/api/ports/:port", handler.GetPort)

	mux.Get("/api/hosts/:host", handler.GetHost)
	mux.Post("/api/hosts", handler.PostHost)

	return mux
}
