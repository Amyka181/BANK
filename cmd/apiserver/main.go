package main

import (
	"Bankirka/infrastructure/postgres"
	"Bankirka/internal/service"
	bankServ "Bankirka/pkg/http"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
)

func main() {

	//c := cache.New()
	//bankService := service.New(c)
	db := postgres.NewDB()
	bankService := service.New(db)

	r := chi.NewRouter()
	r.Use(bankServ.MetricsMiddleware)
	h := bankServ.NewBankHandler(bankService)
	h.ApiRoute(r)
	r.Mount("/debug/pprof/", http.StripPrefix("/debug/pprof", http.HandlerFunc(pprof.Index)))

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Panic("Сервер недоступен")
	}

}
