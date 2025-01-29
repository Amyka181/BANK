package http

import (
	"Bankirka/internal/entity"
	"encoding/json"
	"io"
	"net/http"
)

func (h *BankHandler) ChangeBalanceHandler(w http.ResponseWriter, req *http.Request) {
	body := req.Body
	byteBody, err := io.ReadAll(body)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	defer body.Close()

	var userReq userRequestChange
	err = json.Unmarshal(byteBody, &userReq)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return

	}

	_, err = h.bankService.ChangeBal(userReq.Operation, entity.Difference{userReq.Quantity}, userReq.ID)
	if err != nil {
		BeautifulErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	OkResponse(w, http.StatusOK, "Успешно")

}
