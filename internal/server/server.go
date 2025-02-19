package server

import (
	"Bankirka/internal/entity"
	"Bankirka/internal/service"
	"Bankirka/pkg/grpc"
	"fmt"
)

type Server struct {
	grpc.UnimplementedUserServiceServer
	BankService service.BankService
}

func NewServer(bs service.BankService) *Server {
	return &Server{BankService: bs}
}

func (s *Server) Show(req *grpc.ShowRequest) (*grpc.ShowResponse, error) {
	fmt.Printf("Получен запрос на пользователя ID: %d\n", req.Id)
	resp, err := s.BankService.Show(entity.User{ID: int(req.Id)})
	if err != nil {
		return nil, err
	}
	return &grpc.ShowResponse{Id: int32(resp.ID), Balance: int32(resp.Balance.Money)}, nil
}

func (s *Server) Create(req *grpc.CreateRequest) (*grpc.CreateResponse, error) {
	fmt.Printf("Получен запрос на пользователя ID: %d\n", req.Id)
	resp, err := s.BankService.CreateUser(int(req.Id), entity.Balance{Money: int(req.Balance)})
	if err != nil {
		return nil, err
	}
	return &grpc.CreateResponse{Id: int32(resp.ID), Balance: int32(resp.Balance.Money)}, nil
}

func (s *Server) Change(req *grpc.ChangeRequest) (*grpc.ChangeResponse, error) {
	fmt.Printf("Получен запрос на пользователя ID: %d\n", req.Id)
	resp, err := s.BankService.ChangeBal(req.Operation, entity.Difference{Quantity: int(req.Quantity)}, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &grpc.ChangeResponse{Id: int32(resp.ID), Balance: int32(resp.Balance.Money)}, nil
}
