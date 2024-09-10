package service

import (
	"Bankirka/internal/entity"
)

func (b *BankService) Add(amount entity.Difference, id int) (*entity.User, error) {
	//_, err := b.db.ShowBalance(id)
	//if err != nil {
	//	return nil, err
	//}

	er := b.db.ChangeBalance(id, amount)
	if er != nil {
		return nil, er
	}
	bal, _ := b.db.ShowBalance(id)
	return &entity.User{ID: id, Balance: entity.Balance{Money: bal}}, nil

}
