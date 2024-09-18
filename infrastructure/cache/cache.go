package cache

import (
	"Bankirka/internal/entity"
	"Bankirka/internal/service"
	"sync"
)

type Bd struct {
	person map[int]entity.Balance
	mu     sync.Mutex
}

func New() *Bd {
	return &Bd{
		person: make(map[int]entity.Balance),
	}
}

func (b *Bd) CreatePerson(id int, bal entity.Balance) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	_, ok := b.person[id]
	if ok {
		return service.AccountExistErr
	}

	b.person[id] = bal
	return nil
}

func (b *Bd) ChangeBalance(id int, dif entity.Difference) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	bal, ok := b.person[id]
	if !ok {
		return service.NoAccountErr
	}

	a := bal.Money + dif.Quantity
	b.person[id] = entity.Balance{a}
	return nil
}

func (b *Bd) ShowBalance(id int) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	_, ok := b.person[id]
	if !ok {
		return 0, service.NoAccountErr
	}

	return b.person[id].Money, nil
}
