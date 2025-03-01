package entity

type User struct {
	ID      int
	Balance Balance
}

type Balance struct {
	Money int
}

type Difference struct {
	Quantity int
}

type UpdateUser struct {
	ID     int
	Change int
}
