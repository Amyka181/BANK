package http

import (
	"Bankirka/internal/entity"
	"encoding/json"
	"io"
	"net/http"
)

func (h *BankHandler) ChangeBalanceHttp(w http.ResponseWriter, req *http.Request) {
	body := req.Body
	byteBody, err := io.ReadAll(body)
	if err != nil {
		return
	}

	var userReq userRequestChange
	err = json.Unmarshal(byteBody, &userReq)
	if err != nil {
		return
	}

	if userReq.Operation == "пополнить" {
		_, err2 := h.bankService.Add(entity.Difference{userReq.Quantity}, userReq.ID)
		if err2 != nil {
			return
		}
	} else if userReq.Operation == "снять" {
		_, err3 := h.bankService.AntiAdd(entity.Difference{userReq.Quantity}, userReq.ID)
		if err3 != nil {
			return
		}
	} else {
		return
	}

}
