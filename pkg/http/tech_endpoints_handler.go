package http

import (
	"net/http"
)

func version(w http.ResponseWriter, req *http.Request) {

	// TODO: Взять из конфига ХЗ ВАЩЕ КАК
	//
	//
	//vers := &config.BDVersion{
	//	Version: os.Getenv("VERSION"),
	//}
	//json.NewEncoder(w).Encode(vers)
}
