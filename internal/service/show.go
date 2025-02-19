package service

import (
	"Bankirka/internal/entity"
)

func (b *BankService) Show(person entity.User) (*entity.User, error) {
	bal, err := b.Db.ShowBalance(person)
	if err != nil {
		return nil, err
	}

	return &entity.User{ID: person.ID, Balance: entity.Balance{Money: bal.Balance.Money}}, nil
}
