package service

import "Bankirka/internal/entity"

type BankInt interface {
	CreatePerson(id int, bal entity.Balance) error
	ChangeBalance(id int, dif entity.Difference) error
	ShowBalance(person entity.User) (*entity.User, error)
}
