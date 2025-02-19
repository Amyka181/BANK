package service

import (
	"Bankirka/internal/entity"
)

func (b *BankService) CreateUser(person entity.User) (*entity.User, error) {
	if person.Balance.Money < 0 {
		return nil, NegativeBalanceErr
	}

	err := b.Db.CreatePerson(person)
	if err != nil {
		return nil, err
	}

	return &entity.User{ID: person.ID, Balance: person.Balance}, nil

}
