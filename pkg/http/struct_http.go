package http

type userRequestCreate struct {
	ID      int `json:"ID"`
	Balance int `json:"Balance"`
}

type userRequestShow struct {
	ID int `json:"ID"`
}

type userRequestChange struct {
	ID        int    `json:"ID"`
	Quantity  int    `json:"quantity"`
	Operation string `json:"Operation"`
}
