package service

import (
	"Bankirka/internal/entity"
)

func (b *BankService) CreateUser(id int, bal entity.Balance) (*entity.User, error) {
	if bal.Money < 0 {
		return nil, NegativeBalanceErr
	}

	err := b.db.CreatePerson(id, bal)
	if err != nil {
		return nil, err
	}

	return &entity.User{ID: id, Balance: bal}, nil

}
