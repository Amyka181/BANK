package cache

import (
	"Bankirka/internal/entity"
	"errors"
	"fmt"
)

var (
	NegativeBalanceErr = errors.New("Ваш баланс отрицательный")
	AccountExistErr    = errors.New("Аккаунт уже существует")
	NoAccountErr       = errors.New("Аккаунта не существует")
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
	if bal.Money < 0 {
		return NegativeBalanceErr
	}

	_, ok := b.person[id]
	if ok {
		return AccountExistErr
	}

	b.person[id] = bal
	return nil
}

func (b *bd) ChangeBalance(id int, dif entity.Difference) error {
	bal, ok := b.person[id]
	if !ok {
		return NoAccountErr
	}

	if bal.Money+dif.Quantity < 0 {
		return NegativeBalanceErr
	}

	a := bal.Money + dif.Quantity
	b.person[id] = entity.Balance{a}
	return nil
}

func (b *bd) ShowBalance(id int) (int, error) {
	_, ok := b.person[id]
	if !ok {
		return 0, NoAccountErr
	}

	fmt.Println("Ваш баланс:", b.person[id].Money)

	return b.person[id].Money, nil
}
