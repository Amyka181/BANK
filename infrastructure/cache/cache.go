package cache

import (
	"Bankirka/internal/entity"
	"errors"
)

var (
	negativeBalanceErr = errors.New("your balance is negative")
	noEnoughMoneyErr   = errors.New("not enough money to receive")
	accountExistErr    = errors.New("account already exists")
	noAccountErr       = errors.New("account does not exist")
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
		return negativeBalanceErr
	}

	_, ok := b.person[id]
	if ok {
		return accountExistErr
	}

	b.person[id] = bal
	return nil
}

func (b *bd) ChangeBalance(id int, dif entity.Difference) error {
	bal, ok := b.person[id]
	if !ok {
		return noAccountErr
	}

	if bal.Money+dif.Quantity < 0 {
		return noEnoughMoneyErr
	}

	a := bal.Money + dif.Quantity
	b.person[id] = entity.Balance{a}
	return nil
}

func (b *bd) ShowBalance(id int) (int, error) {
	_, ok := b.person[id]
	if !ok {
		return 0, noAccountErr
	}

	return b.person[id].Money, nil
}
