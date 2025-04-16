package entity

type User struct {
	ID      int     `json:"id"`
	Balance Balance `json:"balance"`
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
