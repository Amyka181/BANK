package main

import (
	"Bankirka/infrastructure/cache"
	"Bankirka/internal/service"
	"Bankirka/pkg/http"
	"github.com/go-chi/chi/v5"
)

func main() {
	c := cache.New()
	bankService := service.New(c)
	r := chi.NewRouter()
	h := http.NewBankHandler(bankService)
	r.Post("/create", h.CreatePersonHttp)

	http.ListenAndServe("localhost:8080", r)

}
