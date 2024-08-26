package service

import (
	"Bankirka/internal/entity"
	"fmt"
)

func Show(person entity.User) {
	fmt.Println("На балансе:", person.Balance.Money)

}
