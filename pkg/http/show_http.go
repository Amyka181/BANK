package http

import (
	"Bankirka/internal/entity"
	"encoding/json"
	"io"
	"net/http"
)

func (h *BankHandler) ShowBalanceHttp(w http.ResponseWriter, req *http.Request) {
	body := req.Body
	byteBody, err := io.ReadAll(body)
	if err != nil {
		return
	}

	var userReq userRequestShow
	err = json.Unmarshal(byteBody, &userReq)
	if err != nil {
		return
	}

	_, err2 := h.bankService.Show(entity.User{ID: userReq.ID})
	if err2 != nil {
		return
	}

}
