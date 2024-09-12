package service

import "errors"

var (
	NegativeBalanceErr      = errors.New("your balance is negative")
	NoEnoughMoneyErr        = errors.New("not enough money to receive")
	AccountExistErr         = errors.New("account already exists")
	NoAccountErr            = errors.New("account does not exist")
	InvalidOperation        = errors.New("invalid operation")
	NegativeValueBalanceErr = errors.New("negative value in changed balance")
)
