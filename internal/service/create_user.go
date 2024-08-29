package service

import "Bankirka/internal/entity"

func (b *BankService) CreateUser(id int, bal entity.Balance) (*entity.User, error) {
	err := b.db.CreatePerson(id, bal)
	if err != nil {
		return nil, err
	}

	return &entity.User{ID: id, Balance: bal}, nil

}
