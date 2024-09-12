package cache

import (
	"Bankirka/internal/entity"
	"Bankirka/internal/service"
	"sync"
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
		return service.AccountExistErr
	}

	b.person[id] = bal
	return nil
}

func (b *bd) ChangeBalance(id int, dif entity.Difference) error {
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

func (b *bd) ShowBalance(id int) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	_, ok := b.person[id]
	if !ok {
		return 0, service.NoAccountErr
	}

	return b.person[id].Money, nil
}
