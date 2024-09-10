package http

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, stCode int, err error) {
	w.WriteHeader(stCode)
	w.Write([]byte(err.Error()))
}

func OkResponse(w http.ResponseWriter, stCode int, value interface{}) {
	resp, err := json.Marshal(value)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(stCode)
	w.Write(resp)
}

func BeautifulErrorResponse(w http.ResponseWriter, stCode int, err error) {
	w.WriteHeader(stCode)
	respErr, _ := json.Marshal(&errorResponce{
		Error: err.Error(),
	})
	w.Write(respErr)

}
