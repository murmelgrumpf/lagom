package routing

import (
	"lagom/handlers"
	"lagom/handlers/viewhandler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router(public http.HandlerFunc) *chi.Mux {
	router := chi.NewMux()
	router.Handle("/*", public)

	registerViews(router)
	registerComponents(router)
	registerActions(router)

	return router
}

func registerViews(router *chi.Mux) {

	router.Get("/", handlers.Make(viewhandler.HandleHome))
}

func registerComponents(router *chi.Mux) {

}

func registerActions(router *chi.Mux) {

}
