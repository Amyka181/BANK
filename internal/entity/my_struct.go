package entity

// TODO: если структуры отдаются в ответе http то лучше добавить к ним теги в snake_case
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
