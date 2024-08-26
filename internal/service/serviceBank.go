package service

import "Bankirka/internal/entity"

type BankInt interface {
	CreatePerson(id int, bal entity.Balance) error
	ChangeBalance(id int, bal entity.Balance, dif entity.Difference) error
	ShowBalance(id int) (int, error)
}
