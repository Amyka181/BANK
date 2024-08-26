package main

import "fmt"

func main() {
	var m Balance = Balance{1000}
	var Roma User = User{1, "Roma", m}
	var diff Difference

	fmt.Scan(&diff.quantity)
	Add(diff, &Roma)
	fmt.Scan(&diff.quantity)
	AntiAdd(diff, &Roma)
	Show(Roma)
	fmt.Println(Roma)

}
