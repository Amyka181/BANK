package http

import (
	"Bankirka/config"
	"encoding/json"
	"net/http"
	"os"
)

func version(w http.ResponseWriter, req *http.Request) {

	vers := &config.BDVersion{
		Version: os.Getenv("VERSION"),
	}
	json.NewEncoder(w).Encode(vers)
}
