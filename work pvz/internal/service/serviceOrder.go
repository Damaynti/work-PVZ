package service

import (
	"context"
	"errors"

	"example.com/mymodule/internal/model"
	"example.com/mymodule/internal/pkg/db/repository/postgres"
)

type Status int

const (
	Issued Status = iota
	Refund
	Defect
)

func parseStatus(input string) Status {
	switch input {
	case "выдан":
		return Issued
	case "возврат":
		return Refund
	case "брак":
		return Defect
	default:
		return -1
	}
}


type ServiceOrder struct {
	Repo *postgres.OrderRepo
}

func NewOrder(repo *postgres.OrderRepo) *ServiceOrder {
	return &ServiceOrder{Repo: repo}
}

// создание
func (s *ServiceOrder) CreateOrder(ctx context.Context,input model.OrderInput) (int64,error) {

	if len(input.FullName) == 0 {
		return 0,errors.New("пустое ФИО")
	}
	if len(input.OrderCode) == 0 {
		return 0,errors.New("пустой код заказа")
	}

	id, err := s.Repo.Add(ctx, &input)
	return id,err

}

// изменение статуса
func (s *ServiceOrder) StatusOrder(ctx context.Context,orderStatus model.OrderStatus) error {
	if orderStatus.ID == 0 {
		return errors.New("нулевой id")
	}
	if len(orderStatus.Status)==0{
		return errors.New("не указан статус заказа")
	}
	if parseStatus(orderStatus.Status) == -1 {
		return errors.New("статус заказа указан некоректно")
	}
	return s.Repo.StatusOrder(ctx,orderStatus)
}

// удаление
func (s *ServiceOrder) DeleteOrder(ctx context.Context,id int64) error {
	if id == 0 {
		return errors.New("нулевой id")
	}

	err := s.Repo.DeleteOrder(ctx, int64(id))
	return err

}

// поиск заказа
func (s *ServiceOrder) SearchOrder(ctx context.Context,input model.OrderSerch) (int64,error) {
	if len(input.FullName) == 0 {
		return 0,errors.New("пустое ФИО")
	}
	if len(input.OrderCode) == 0 {
		return 0,errors.New("пустой код заказа")
	}
	id,err:=s.Repo.SearchOrder(ctx,input)
	return id,err
}

// лист
func (s *ServiceOrder) ListOrder(ctx context.Context) ([]model.Order, error) {
	orders, err := s.Repo.GetAllOrder(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
