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

func (h *BankHandler) CreatePersonHttp(w http.ResponseWriter, req *http.Request) {
	body := req.Body

	byteBody, err := io.ReadAll(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

	}
	var userReq userRequestCreate
	err = json.Unmarshal(byteBody, &userReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

	}

	user, err := h.bankService.CreateUser(userReq.ID, entity.Balance{Money: userReq.Balance})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respErr, err := json.Marshal(&errorResponce{
			Error: err.Error(),
		})

		if err != nil {
			w.Write([]byte(err.Error()))
		}
		w.Write(respErr)

	} else {
		responce, err := json.Marshal(user)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(responce)

	}
}
