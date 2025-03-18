package service

import (
	"Bankirka/internal/entity"
	"time"
)

func (b *BankService) add(amount entity.Difference, id int) (*entity.User, error) {
	if amount.Quantity < 0 {
		return nil, NegativeValueBalanceErr
	}

	//var userUpdate *entity.UpdateUser
	//userUpdate.ID = id
	//userUpdate.Change = amount.Quantity
	userUpdate := entity.UpdateUser{ID: id, Change: amount.Quantity}

	SendToRabbit(userUpdate)
	time.Sleep(1 * time.Second)

	//err := b.Db.ChangeBalance(id, amount)
	//if err != nil {
	//	return nil, err
	//}
	bal, _ := b.Db.ShowBalance(entity.User{ID: id})
	return &entity.User{ID: id, Balance: entity.Balance{Money: bal.Balance.Money}}, nil

}
