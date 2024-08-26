package cache

import (
	"Bankirka/internal/entity"
	"fmt"
)

type bd struct {
	person map[int]entity.Balance
}

func New() *bd {
	return &bd{
		person: make(map[int]entity.Balance),
	}
}

func (b *bd) CreatePerson(id int, bal entity.Balance) error {
	b.person[id] = bal
}

func (b *bd) ChangeBalance(id int, bal entity.Balance, dif entity.Difference) error {
	a := bal.Money + dif.Quantity
	b.person[id] = entity.Balance{a}
}

func (b *bd) ShowBalance(id int) (int, error) {
	fmt.Println("Ваш баланс:", b.person[id])
}
