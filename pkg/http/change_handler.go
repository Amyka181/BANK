package http

import (
	"Bankirka/internal/entity"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var invalidOperation = errors.New("invalid operation")

func (h *BankHandler) ChangeBalanceHandler(w http.ResponseWriter, req *http.Request) {
	body := req.Body
	byteBody, err := io.ReadAll(body)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
	}
	defer body.Close()

	var userReq userRequestChange
	err = json.Unmarshal(byteBody, &userReq)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)

	}

	switch userReq.Operation {
	case "пополнить":
		user, err := h.bankService.Add(entity.Difference{userReq.Quantity}, userReq.ID)
		if err != nil {
			BeautifulErrorResponse(w, http.StatusBadRequest, err)
		} else {
			OkResponse(w, http.StatusOK, user)
		}
	case "снять":
		user, err := h.bankService.AntiAdd(entity.Difference{userReq.Quantity}, userReq.ID)
		if err != nil {
			BeautifulErrorResponse(w, http.StatusBadRequest, err)
		} else {
			OkResponse(w, http.StatusOK, user)
		}

	default:
		BeautifulErrorResponse(w, http.StatusBadRequest, invalidOperation)
	}

}
