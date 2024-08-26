package service

import "Bankirka/internal/entity"

type BankInt interface {
	CreatePerson(id int, bal entity.Balance)
	ChangeBalance(id int, bal entity.Balance, dif entity.Difference)
	ShowBalance(id int)
}
