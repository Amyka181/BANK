package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (h *BankHandler) ApiRoute(r *chi.Mux) {

	r.Post("/create", h.CreatePersonHandler)
	r.Post("/change", h.ChangeBalanceHandler)
	r.Post("/show", h.ShowBalanceHandler)
	r.Get("/version", version)
	r.Handle("/metrics", promhttp.Handler())
}
