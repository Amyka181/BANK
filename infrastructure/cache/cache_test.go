package cache

import (
	"Bankirka/internal/entity"
	"Bankirka/internal/service"
	mock_service "Bankirka/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

//func TestNew(t *testing.T) {
//	tests := []struct {
//		name string
//		want *bd
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := New(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("New() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func NewMockService(mockRep *mock_service.MockBankInt) *service.BankService {
	return &service.BankService{Db: mockRep}
}

func Test_bd_CreatePerson(t *testing.T) {

	type args struct {
		id  int
		bal entity.Balance
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRep := mock_service.NewMockBankInt(ctrl)
	mockService := NewMockService(mockRep)

	tests := []struct {
		name    string
		args    args
		prepare func()
		want    *entity.User
		wantErr error
	}{
		{
			name: "Create person success",
			args: args{id: 1, bal: entity.Balance{Money: 1000}},
			prepare: func() {
				mockRep.EXPECT().CreatePerson(1, entity.Balance{Money: 1000}).Return(nil)
			},
			want:    &entity.User{ID: 1, Balance: entity.Balance{Money: 1000}},
			wantErr: nil,
		},
		{
			name: "Create person failed, account already exists",
			args: args{id: 1, bal: entity.Balance{Money: 100}},
			prepare: func() {
				mockRep.EXPECT().CreatePerson(1, entity.Balance{Money: 100}).Return(service.AccountExistErr)
			},
			want:    nil,
			wantErr: service.AccountExistErr,
		},
		{
			name: "Create person failed, balance is negative",
			args: args{id: 2, bal: entity.Balance{Money: -300}},
			prepare: func() {

			},
			want:    nil,
			wantErr: service.NegativeBalanceErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			got, err := mockService.CreateUser(tt.args.id, tt.args.bal)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_bd_ChangeBalance(t *testing.T) {
	type args struct {
		operation string
		amount    entity.Difference
		id        int
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRep := mock_service.NewMockBankInt(ctrl)
	mockService := NewMockService(mockRep)

	tests := []struct {
		name    string
		args    args
		prepare func()
		want    *entity.User
		wantErr error
	}{
		{
			name: "Change balance success (add)",
			args: args{operation: "пополнить", amount: entity.Difference{Quantity: 500}, id: 1},
			prepare: func() {
				mockRep.EXPECT().ChangeBalance(1, entity.Difference{Quantity: 500}).Return(nil)
				mockRep.EXPECT().ShowBalance(1).Return(1500, nil)
			},
			want:    &entity.User{ID: 1, Balance: entity.Balance{Money: 1500}},
			wantErr: nil,
		},
		{
			name: "Change balance success (take)",
			args: args{operation: "снять", amount: entity.Difference{Quantity: 500}, id: 1},
			prepare: func() {
				mockRep.EXPECT().ShowBalance(1).Return(1500, nil)
				mockRep.EXPECT().ChangeBalance(1, entity.Difference{Quantity: -500}).Return(nil)
				mockRep.EXPECT().ShowBalance(1).Return(1000, nil)
			},
			want:    &entity.User{ID: 1, Balance: entity.Balance{Money: 1000}},
			wantErr: nil,
		},
		{
			name: "Change balance failed, account does not exist",
			args: args{operation: "пополнить", amount: entity.Difference{Quantity: 1000}, id: 10},
			prepare: func() {
				mockRep.EXPECT().ChangeBalance(10, entity.Difference{Quantity: 1000}).Return(service.NoAccountErr)
			},
			want:    nil,
			wantErr: service.NoAccountErr,
		},
		{
			name: "Change balance failed, not enough money to receive",
			args: args{operation: "снять", amount: entity.Difference{Quantity: 5000}, id: 1},
			prepare: func() {
				mockRep.EXPECT().ShowBalance(1).Return(1000, nil)
			},
			want:    nil,
			wantErr: service.NoEnoughMoneyErr,
		},
		{
			name: "Change balance failed, invalid operation",
			args: args{operation: "положить", amount: entity.Difference{Quantity: 1000}, id: 1},
			prepare: func() {

			},
			want:    nil,
			wantErr: service.InvalidOperation,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			got, err := mockService.ChangeBal(tt.args.operation, tt.args.amount, tt.args.id)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_bd_ShowBalance(t *testing.T) {
	type args struct {
		person entity.User
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRep := mock_service.NewMockBankInt(ctrl)
	mockService := NewMockService(mockRep)

	tests := []struct {
		name    string
		args    args
		prepare func()
		want    *entity.User
		wantErr error
	}{
		{
			name: "Show balance success",
			args: args{person: entity.User{ID: 1, Balance: entity.Balance{Money: 1000}}},
			prepare: func() {
				mockRep.EXPECT().ShowBalance(1).Return(1000, nil)
			},
			want:    &entity.User{ID: 1, Balance: entity.Balance{Money: 1000}},
			wantErr: nil,
		},
		{
			name: "Show balance failed, account does not exist",
			args: args{person: entity.User{ID: 7, Balance: entity.Balance{Money: 800}}},
			prepare: func() {
				mockRep.EXPECT().ShowBalance(7).Return(0, service.NoAccountErr)
			},
			want:    nil,
			wantErr: service.NoAccountErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			got, err := mockService.Show(tt.args.person)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
