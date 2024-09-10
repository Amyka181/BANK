package main

import (
	"Bankirka/infrastructure/cache"
	"Bankirka/internal/service"
	bankServ "Bankirka/pkg/http"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	c := cache.New()
	bankService := service.New(c)
	r := chi.NewRouter()
	h := bankServ.NewBankHandler(bankService)
	r.Post("/create", h.CreatePersonHandler)
	r.Post("/change", h.ChangeBalanceHandler)
	r.Post("/show", h.ShowBalanceHandler)

	http.ListenAndServe("localhost:8080", r)

}
