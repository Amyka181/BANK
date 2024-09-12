package service

import (
	"Bankirka/internal/entity"
)

func (b *BankService) ChangeBal(operation string, amount entity.Difference, id int) (*entity.User, error) {
	switch operation {
	case "пополнить":
		user, err := b.add(amount, id)
		if err != nil {
			return nil, err
		} else {
			return user, nil
		}
	case "снять":
		user, err := b.antiAdd(amount, id)
		if err != nil {
			return nil, err
		} else {
			return user, nil
		}
	default:
		return nil, InvalidOperation
	}
}