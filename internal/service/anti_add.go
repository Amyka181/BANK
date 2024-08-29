package service

import (
	"Bankirka/internal/entity"
)

func (b *BankService) AntiAdd(amount entity.Difference, id int) (*entity.User, error) {
	_, err := b.db.ShowBalance(id)
	if err != nil {
		return nil, err
	}
	amount.Quantity = amount.Quantity * (-1)
	er := b.db.ChangeBalance(id, amount)
	if er != nil {
		return nil, er
	}
	bal, _ := b.db.ShowBalance(id)

	return &entity.User{ID: id, Balance: entity.Balance{Money: bal}}, nil
}
