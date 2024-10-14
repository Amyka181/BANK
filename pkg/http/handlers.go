package http

import "github.com/go-chi/chi/v5"

func (h *BankHandler) ApiRoute(r *chi.Mux) {

	r.Post("/create", h.CreatePersonHandler)
	r.Post("/change", h.ChangeBalanceHandler)
	r.Post("/show", h.ShowBalanceHandler)

}
