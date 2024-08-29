package http

import (
	"Bankirka/internal/entity"
	"encoding/json"
	"io"
	"net/http"
)

func (h *BankHandler) CreatePersonHttp(w http.ResponseWriter, req *http.Request) {
	body := req.Body

	byteBody, err := io.ReadAll(body)
	if err != nil {
		return
	}
	var userReq userRequestCreate
	err = json.Unmarshal(byteBody, &userReq)
	if err != nil {
		return
	}

	_, err2 := h.bankService.CreateUser(userReq.ID, entity.Balance{Money: userReq.Balance})
	if err2 != nil {
		return
	}

	w.Write([]byte("Аккаунт создан"))
}
