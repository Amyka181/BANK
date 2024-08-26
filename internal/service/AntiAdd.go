package service

import (
	"Bankirka/internal/entity"
	"fmt"
)

func (b *bankService) AntiAdd(amount entity.Difference, id int) (*entity.User, error) {
	currentBalance, err := b.db.ShowBalance(id)
	if err != nil {
		return
	}
	b.db.ChangeBalance()
	person.Balance.Money = person.Balance.Money - amount.Quantity
	fmt.Println(person.ID, "снял с баланcа сумму:", amount.Quantity)
}
