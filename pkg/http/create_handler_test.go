package http

import (
	"Bankirka/infrastructure/cache"
	"Bankirka/internal/entity"
	"Bankirka/internal/service"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	c           *cache.Bd
	bankService *service.BankService
	h           *BankHandler
	handler     http.HandlerFunc
)

func TestMain(m *testing.M) {
	c = cache.New()
	bankService = service.New(c)
	h = NewBankHandler(bankService)
	m.Run()
}

func TestBankHandler_CreatePersonHandler(t *testing.T) {

	tests := []struct {
		name    string
		body    userRequestCreate
		want    entity.User
		wantErr errorResponce
	}{
		{
			name: "Create person success",
			body: userRequestCreate{ID: 1, Balance: 1000},
			want: entity.User{ID: 1, Balance: entity.Balance{Money: 1000}},
		},
		{
			name:    "Create person failed, balance is negative",
			body:    userRequestCreate{ID: 2, Balance: -1000},
			wantErr: errorResponce{service.NegativeBalanceErr.Error()},
		},
		{
			name:    "Create person failed, account already exists",
			body:    userRequestCreate{ID: 1, Balance: 1000},
			wantErr: errorResponce{service.AccountExistErr.Error()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler = h.CreatePersonHandler
			bodyReq, err := json.Marshal(tt.body)
			if err != nil {
				return
			}
			req, err := http.NewRequest("POST", "/create", bytes.NewReader(bodyReq))
			if err != nil {
				return
			}
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			res := rr.Result()
			sc := res.StatusCode
			defer res.Body.Close()
			byteBody, _ := io.ReadAll(res.Body)

			if sc == http.StatusCreated {
				var user entity.User
				_ = json.Unmarshal(byteBody, &user)
				assert.Equal(t, tt.want, user)
			} else {
				var respErr errorResponce
				_ = json.Unmarshal(byteBody, &respErr)
				assert.Equal(t, tt.wantErr, respErr)
			}

		})

	}
}

func TestBankHandler_ChangeBalanceHandler(t *testing.T) {

	tests := []struct {
		name    string
		body    userRequestChange
		want    entity.User
		wantErr errorResponce
	}{
		{
			name: "Change balance success (add)",
			body: userRequestChange{ID: 1, Quantity: 500, Operation: "пополнить"},
			want: entity.User{ID: 1, Balance: entity.Balance{Money: 1500}},
		},
		{
			name: "Change balance success (take)",
			body: userRequestChange{ID: 1, Quantity: 500, Operation: "снять"},
			want: entity.User{ID: 1, Balance: entity.Balance{Money: 1000}},
		},
		{
			name:    "Change balance failed, account does not exist",
			body:    userRequestChange{ID: 4, Quantity: 1000, Operation: "пополнить"},
			wantErr: errorResponce{service.NoAccountErr.Error()},
		},
		{
			name:    "Change balance failed, not enough money to receive",
			body:    userRequestChange{ID: 1, Quantity: 5000, Operation: "снять"},
			wantErr: errorResponce{service.NoEnoughMoneyErr.Error()},
		},
		{
			name:    "Change balance failed, invalid operation",
			body:    userRequestChange{ID: 1, Quantity: 1000, Operation: "gjgjkybnm"},
			wantErr: errorResponce{service.InvalidOperation.Error()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler = h.ChangeBalanceHandler
			bodyReq, err := json.Marshal(tt.body)
			if err != nil {
				return
			}
			req, err := http.NewRequest("POST", "/change", bytes.NewReader(bodyReq))
			if err != nil {
				return
			}
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			res := rr.Result()
			sc := res.StatusCode
			defer res.Body.Close()
			byteBody, _ := io.ReadAll(res.Body)

			if sc == http.StatusOK {
				var user entity.User
				_ = json.Unmarshal(byteBody, &user)
				assert.Equal(t, tt.want, user)
			} else {
				var respErr errorResponce
				_ = json.Unmarshal(byteBody, &respErr)
				assert.Equal(t, tt.wantErr, respErr)
			}
		})
	}
}

func TestBankHandler_ShowBalanceHandler(t *testing.T) {

	tests := []struct {
		name    string
		body    userRequestShow
		want    entity.User
		wantErr errorResponce
	}{
		{
			name: "Show balance success",
			body: userRequestShow{ID: 1},
			want: entity.User{ID: 1, Balance: entity.Balance{Money: 1000}},
		},
		{
			name:    "Show balance failed, account does not exist",
			body:    userRequestShow{ID: 7},
			wantErr: errorResponce{service.NoAccountErr.Error()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler = h.ShowBalanceHandler
			bodyReq, err := json.Marshal(tt.body)
			if err != nil {
				return
			}
			req, err := http.NewRequest("POST", "/show", bytes.NewReader(bodyReq))
			if err != nil {
				return
			}
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			res := rr.Result()
			sc := res.StatusCode
			defer res.Body.Close()
			byteBody, _ := io.ReadAll(res.Body)

			if sc == http.StatusOK {
				var user entity.User
				_ = json.Unmarshal(byteBody, &user)
				assert.Equal(t, tt.want, user)
			} else {
				var respErr errorResponce
				_ = json.Unmarshal(byteBody, &respErr)
				assert.Equal(t, tt.wantErr, respErr)
			}
		})
	}
}
