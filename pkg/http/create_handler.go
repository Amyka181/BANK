package http

import (
	"Bankirka/infrastructure/cache"
	"Bankirka/internal/entity"
	"encoding/json"
	"io"
	"net/http"
)

type errorResponce struct {
	Error string
}

func (h *BankHandler) CreatePersonHandler(w http.ResponseWriter, req *http.Request) {
	body := req.Body

	byteBody, err := io.ReadAll(body)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
	}
	defer body.Close()

	var userReq userRequestCreate
	err = json.Unmarshal(byteBody, &userReq)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
	}

	if userReq.Balance < 0 {
		BeautifulErrorResponse(w, http.StatusBadRequest, cache.NegativeBalanceErr)
		//ErrorResponse(w, http.StatusBadRequest, cache.NegativeBalanceErr)
		return
	}

	user, err := h.bankService.CreateUser(userReq.ID, entity.Balance{Money: userReq.Balance})

	if err != nil {
		BeautifulErrorResponse(w, http.StatusBadRequest, err)
	} else {
		OkResponse(w, http.StatusCreated, user)
	}
}
