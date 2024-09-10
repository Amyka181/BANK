package http

import (
	"Bankirka/internal/entity"
	"encoding/json"
	"io"
	"net/http"
)

func (h *BankHandler) ShowBalanceHandler(w http.ResponseWriter, req *http.Request) {
	body := req.Body
	byteBody, err := io.ReadAll(body)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
	}
	defer body.Close()

	var userReq userRequestShow
	err = json.Unmarshal(byteBody, &userReq)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
	}

	user, err := h.bankService.Show(entity.User{ID: userReq.ID})
	if err != nil {
		BeautifulErrorResponse(w, http.StatusBadRequest, err)
	} else {
		OkResponse(w, http.StatusOK, user)
	}
}
