package service

import (
	"Bankirka/internal/entity"
)

func (b *BankService) antiAdd(amount entity.Difference, id int) (*entity.User, error) {
	if amount.Quantity < 0 {
		return nil, InvalidOperation
	}

	money, err := b.Db.ShowBalance(id)
	if err != nil {
		return nil, err
	}

	if money+amount.Quantity*(-1) < 0 {
		return nil, NoEnoughMoneyErr
	}

	amount.Quantity = amount.Quantity * (-1)

	UserUpdate := entity.UpdateUser{ID: id, Change: amount.Quantity}
	SendToRabbit(UserUpdate)

	//er := b.Db.ChangeBalance(id, amount)
	//if er != nil {
	//	return nil, er
	//}
	bal, _ := b.Db.ShowBalance(id)

	return &entity.User{ID: id, Balance: entity.Balance{Money: bal}}, nil
}
