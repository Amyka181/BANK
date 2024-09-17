package service

type BankService struct {
	Db BankInt
}

func New(db BankInt) *BankService {
	return &BankService{Db: db}
}
