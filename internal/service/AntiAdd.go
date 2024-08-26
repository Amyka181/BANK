package service

import (
	"Bankirka/internal/entity"
	"fmt"
)

func AntiAdd(amount entity.Difference, person *entity.User) {
	person.Balance.Money = person.Balance.Money - amount.Quantity
	fmt.Println(person.ID, "снял с баланcа сумму:", amount.Quantity)
}
