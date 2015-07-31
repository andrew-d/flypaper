package router

import (
	"github.com/zenazn/goji/web"

	"github.com/andrew-d/flypaper/handler"
)

func New() *web.Mux {
	mux := web.New()

	mux.Get("/api/port/:port", handler.GetPort)

	return mux
}
