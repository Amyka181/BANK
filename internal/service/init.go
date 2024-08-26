package service

type bankService struct {
	db BankInt
}

func New(db BankInt) *bankService {
	return &bankService{db: db}
}
