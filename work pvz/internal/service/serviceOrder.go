package service

import (
	"errors"

	"example.com/mymodule/internal/model"
)

type storageOrder interface {
	CreateOrder(order model.OrderInput, chOrder chan bool) error
	SearchOrder(input model.OrderInput) error
	StatusOrder(orderStatus model.OrderStatus, chOrder chan bool) error
	DeleteOrder(id int, chOrder chan bool) error
	ListOrder() ([]model.Order, error)
}

type ServiceOrder struct {
	service storageOrder
}

func NewOrder(service storageOrder) ServiceOrder {
	return ServiceOrder{service: service}
}

// создание
func (service ServiceOrder) CreateOrder(input model.OrderInput, chOrder chan bool) error {

	if len(input.FullName) == 0 {
		return errors.New("пустое ФИО")
	}
	if len(input.OrderCode) == 0 {
		return errors.New("пустой код заказа")
	}

	return service.service.CreateOrder(input, chOrder)
}

// изменение статуса
func (service ServiceOrder) StatusOrder(orderStatus model.OrderStatus, chOrder chan bool) error {
	if orderStatus.ID == 0 {
		return errors.New("нулевой id")
	}
	return service.service.StatusOrder(orderStatus, chOrder)
}

// удаление
func (service ServiceOrder) DeleteOrder(id int, chOrder chan bool) error {

	if id == 0 {
		return errors.New("нулевой id")
	}

	return service.service.DeleteOrder(id, chOrder)

}

// поиск заказа
func (service ServiceOrder) SearchOrder(input model.OrderInput) error {
	if len(input.FullName) == 0 {
		return errors.New("пустое ФИО")
	}
	if len(input.OrderCode) == 0 {
		return errors.New("пустой код заказа")
	}

	return service.service.SearchOrder(input)
}

// лист
func (service ServiceOrder) ListOrder() ([]model.Order, error) {
	orders, err := service.service.ListOrder()
	if err != nil {
		return nil, err
	}
	return orders, nil
}
