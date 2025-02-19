package http

import (
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
		return
	}
	defer body.Close()

	var userReq userRequestCreate
	err = json.Unmarshal(byteBody, &userReq)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.bankService.CreateUser(entity.User{ID: userReq.ID, Balance: entity.Balance{userReq.Balance}})

	if err != nil {
		BeautifulErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	OkResponse(w, http.StatusCreated, user)

}
