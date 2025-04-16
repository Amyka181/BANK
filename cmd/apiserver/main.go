package main

import (
	"Bankirka/config"
	"Bankirka/infrastructure/postgres"
	"Bankirka/internal/server"
	"Bankirka/internal/service"
	bankServ "Bankirka/pkg/http"
	"Bankirka/pkg/proto"
	"fmt"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println("Ошибка загрузки конфигурации: ", err)
		return
		//TODO: а попробуй запустить сервер без конфига - ну попоробовала
	}

	//c := cache.New()
	//bankService := service.New(c)
	db := postgres.NewDB(*cfg)
	bankService := service.New(db)
	// HTTP
	r := chi.NewRouter()
	r.Use(bankServ.MetricsMiddleware)
	h := bankServ.NewBankHandler(bankService)
	h.ApiRoute(r)
	r.Mount("/debug/pprof/", http.StripPrefix("/debug/pprof", http.HandlerFunc(pprof.Index)))
	log.Println("Сервер запущен")
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), r)
	if err != nil {
		log.Println("Сервер недоступен")
		return
	}

	// GRPC
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Grpc.GRPCPort))
	if err != nil {
		log.Fatalf("Ошибка при создании слушателя: %v", err)
	}

	grpcServer := grpc.NewServer()
	UserServiceServer := server.NewServer(bankService)
	proto.RegisterUserServiceServer(grpcServer, UserServiceServer)

	log.Println("gRPC сервер запущен на порту 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}

}
