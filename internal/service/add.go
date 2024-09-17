package service

import (
	"Bankirka/internal/entity"
)

func (b *BankService) add(amount entity.Difference, id int) (*entity.User, error) {
	if amount.Quantity < 0 {
		return nil, NegativeValueBalanceErr
	}

	err := b.Db.ChangeBalance(id, amount)
	if err != nil {
		return nil, err
	}
	bal, _ := b.Db.ShowBalance(id)
	return &entity.User{ID: id, Balance: entity.Balance{Money: bal}}, nil

}
