package http

import (
	"Bankirka/internal/entity"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var invalidOperation = errors.New("invalid operation")

func (h *BankHandler) ChangeBalanceHttp(w http.ResponseWriter, req *http.Request) {
	body := req.Body
	byteBody, err := io.ReadAll(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	var userReq userRequestChange
	err = json.Unmarshal(byteBody, &userReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	switch userReq.Operation {
	case "пополнить":
		user, err := h.bankService.Add(entity.Difference{userReq.Quantity}, userReq.ID)
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
			w.WriteHeader(http.StatusOK)
			w.Write(responce)

		}
	case "снять":
		user, err := h.bankService.AntiAdd(entity.Difference{userReq.Quantity}, userReq.ID)
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
			w.WriteHeader(http.StatusOK)
			w.Write(responce)

		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		respErr, err := json.Marshal(&errorResponce{
			Error: invalidOperation.Error(),
		})

		if err != nil {
			w.Write([]byte(err.Error()))
		}
		w.Write(respErr)

	}

}
