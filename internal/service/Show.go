package service

import (
	"Bankirka/internal/entity"
	"fmt"
)

func (b *bankService) Show(person entity.User) {
	fmt.Println("На балансе:", person.Balance.Money)

}
