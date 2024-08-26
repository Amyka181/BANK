package cache

import (
	"Bankirka/internal/entity"
	"fmt"
)

type bd struct {
	person map[int]entity.Balance
}

func (b *bd) CreatePerson(id int, bal entity.Balance) {
	b.person[id] = bal
}

func (b *bd) ChangeBalance(id int, bal entity.Balance, dif entity.Difference) {
	a := bal.Money + dif.Quantity
	b.person[id] = entity.Balance{a}
}

func (b *bd) ShowBalance(id int) {
	fmt.Println("Ваш баланс:", b.person[id])
}
