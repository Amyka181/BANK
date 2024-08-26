package main

import (
	"Bankirka/infrastructure/cache"
	"Bankirka/internal/service"
)

func main() {
	c := cache.New()
	bankService := service.New(c)

	bankService.
}
