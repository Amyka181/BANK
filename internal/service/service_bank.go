package service

import "Bankirka/internal/entity"

type BankInt interface {
	CreatePerson(person entity.User) error
	ChangeBalance(id int, dif entity.Difference) error
	ShowBalance(person entity.User) (*entity.User, error)
}
