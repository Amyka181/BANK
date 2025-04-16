package server

import (
	"Bankirka/internal/entity"
	"Bankirka/pkg/proto"
	"context"
	"fmt"
)

type Server struct {
	proto.UnimplementedUserServiceServer
	BankService BankServiceInt
}
type BankServiceInt interface {
	Add(amount entity.Difference, id int) (*entity.User, error)
	AntiAdd(amount entity.Difference, id int) (*entity.User, error)
	Show(person entity.User) (*entity.User, error)
	CreateUser(person entity.User) (*entity.User, error)
	ChangeBal(operation string, amount entity.Difference, id int) (*entity.User, error)
}

func NewServer(bs BankServiceInt) *Server {
	return &Server{BankService: bs}
}

func (s *Server) Show(ctx context.Context, req *proto.ShowRequest) (*proto.ShowResponse, error) {
	fmt.Printf("Получен запрос на пользователя ID: %d\n", req.Id)
	resp, err := s.BankService.Show(entity.User{ID: int(req.Id)})
	if err != nil {
		return nil, err
	}
	return &proto.ShowResponse{Id: int32(resp.ID), Balance: int32(resp.Balance.Money)}, nil

}

func (s *Server) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	fmt.Printf("Получен запрос на пользователя ID: %d\n", req.Id)
	resp, err := s.BankService.CreateUser(entity.User{ID: int(req.Id), Balance: entity.Balance{Money: int(req.Balance)}})
	if err != nil {
		return nil, err
	}
	return &proto.CreateResponse{Id: int32(resp.ID), Balance: int32(resp.Balance.Money)}, nil
}

func (s *Server) Change(ctx context.Context, req *proto.ChangeRequest) (*proto.ChangeResponse, error) {
	fmt.Printf("Получен запрос на пользователя ID: %d\n", req.Id)
	resp, err := s.BankService.ChangeBal(req.Operation, entity.Difference{Quantity: int(req.Quantity)}, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &proto.ChangeResponse{Id: int32(resp.ID), Balance: int32(resp.Balance.Money)}, nil
}
