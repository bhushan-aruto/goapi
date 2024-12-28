package internal

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(middleware.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Logger)
		router.Get("/coins", GetCoinBalence)
	})
}
