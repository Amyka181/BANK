package main

import (
	"Bankirka/infrastructure/postgres"
	"Bankirka/internal/service"
	bankServ "Bankirka/pkg/http"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {

	//c := cache.New()
	//bankService := service.New(c)
	db := postgres.NewDB()
	bankService := service.New(db)

	r := chi.NewRouter()
	h := bankServ.NewBankHandler(bankService)
	h.ApiRoute(r)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Panic("Сервер недоступен")
	}

}
