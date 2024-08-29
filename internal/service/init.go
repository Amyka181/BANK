package service

type BankService struct {
	db BankInt
}

func New(db BankInt) *BankService {
	return &BankService{db: db}
}
