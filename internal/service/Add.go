package service

import (
	"Bankirka/internal/entity"
	"fmt"
)

func (b *bankService) Add(amount entity.Difference, person *entity.User) {
	person.Balance.Money = person.Balance.Money + amount.Quantity
	fmt.Println(person.ID, "пополнил баланc на сумму:", amount.Quantity)
}
