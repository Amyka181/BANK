package http

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type Version struct {
	Version string
}

func version(w http.ResponseWriter, req *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла: %v", err)
	}

	vers := &Version{
		Version: os.Getenv("VERSION"),
	}
	json.NewEncoder(w).Encode(vers)
}
