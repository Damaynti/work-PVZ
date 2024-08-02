package service

import (
	"errors"

	"example.com/mymodule/internal/model"
)

type storage interface {
	Create(order model.OrderInput)error
	Search(input model.OrderInput)error
	Status(orderStatus model.OrderStatus)error
	Del(id int)error
	List()([]model.Order,error)
}

type Service struct{
	s storage
}
func New (s storage) Service{
	return Service{s:s}
}

// создание 
func (s Service) Create(input model.OrderInput) error{
	
	if len(input.FullName)==0{
		return errors.New("пустое ФИО")
	}
	if len(input.OrderCode)==0{
		return errors.New("пустой код заказа")
	}

	return s.s.Create(input)
}

// изменение статуса
func(s Service) Status(orderStatus model.OrderStatus)error{
	if orderStatus.ID==0{
		return errors.New("нулевой id")
	}
	return s.s.Status(orderStatus)
}

// удаление
func (s Service) Del (id int) error{
	
	if id==0{
		return errors.New("нулевой id")
	}

	return s.s.Del(id)
}

// поиск заказа
func (s Service) Search(input model.OrderInput)error{
	if len(input.FullName)==0{
		return errors.New("пустое ФИО")
	}
	if len(input.OrderCode)==0{
		return errors.New("пустой код заказа")
	}

	return s.s.Search(input)
}

// печать 
func (s Service) List() ([]model.Order, error){
	orders,err :=s.s.List()
	if err!=nil{
		return nil, err
	}
	return orders,nil
}