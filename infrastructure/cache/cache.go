package cache

import (
	"Bankirka/internal/entity"
	"errors"
	"sync"
)

var (
	NegativeBalanceErr = errors.New("your balance is negative")
	NoEnoughMoneyErr   = errors.New("not enough money to receive")
	AccountExistErr    = errors.New("account already exists")
	NoAccountErr       = errors.New("account does not exist")
)

type bd struct {
	person map[int]entity.Balance
	mu     sync.Mutex
}

func New() *bd {
	return &bd{
		person: make(map[int]entity.Balance),
	}
}

func (b *bd) CreatePerson(id int, bal entity.Balance) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	_, ok := b.person[id]
	if ok {
		return AccountExistErr
	}

	b.person[id] = bal
	return nil
}

func (b *bd) ChangeBalance(id int, dif entity.Difference) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	bal, ok := b.person[id]
	if !ok {
		return NoAccountErr
	}

	a := bal.Money + dif.Quantity
	b.person[id] = entity.Balance{a}
	return nil
}

func (b *bd) ShowBalance(id int) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	_, ok := b.person[id]
	if !ok {
		return 0, NoAccountErr
	}

	return b.person[id].Money, nil
}
