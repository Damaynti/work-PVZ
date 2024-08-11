package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"example.com/mymodule/internal/model"
	"example.com/mymodule/internal/service"
	"example.com/mymodule/internal/storage"
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

var muPVZ sync.Mutex
var muOrder sync.Mutex

func OrderCreate(list []string) {
	if len(list) != 3 {
		fmt.Println("Неправильно введена команда")
		return
	}
	muOrder.Lock()
	defer muOrder.Unlock()
	defer wg.Done()
	storOrder, err := storage.NewOrder()
	if err != nil {
		fmt.Println("ошибка при создании хранилища:", err)
		return
	}

	defer storOrder.Close()
	serviceOrder := service.NewOrder(&storOrder)
	fullname := list[1]
	if fullname == "" {
		log.Println("не указано ФИО")
		return
	}
	code := list[2]
	if code == "" {
		log.Println("не указан код товара")
		return
	}

	err = serviceOrder.CreateOrder(model.OrderInput{FullName: fullname, OrderCode: code}, chOrder)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("заказ добавлен")
}

func OrderDelete(list []string) {
	if len(list) != 2 {
		fmt.Println("Неправильно введена команда")
		return
	}
	muOrder.Lock()
	defer muOrder.Unlock()
	defer wg.Done()
	storOrder, err := storage.NewOrder()
	if err != nil {
		fmt.Println("ошибка при создании хранилища:", err)
		return
	}
	defer storOrder.Close()
	serviceOrder := service.NewOrder(&storOrder)
	id, err := strconv.Atoi(list[1])
	if err != nil {
		fmt.Println("неверно введено id", err)
		return
	}

	if id == 0 {
		log.Println("не указано id заказа")
		return
	}
	err = serviceOrder.DeleteOrder(id, chOrder)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("заказ удален")
}

func OrderList() {
	muOrder.Lock()
	defer muOrder.Unlock()
	defer wg.Done()
	storOrder, err := storage.NewOrder()
	if err != nil {
		fmt.Println("ошибка при создании хранилища:", err)
		return
	}
	defer storOrder.Close()
	serviceOrder := service.NewOrder(&storOrder)
	list, err := serviceOrder.ListOrder()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", list)
}

func OrderStatus(list []string) {
	if len(list) != 3 {
		fmt.Println("Неправильно введена команда")
		return
	}
	muOrder.Lock()
	defer muOrder.Unlock()
	defer wg.Done()
	storOrder, err := storage.NewOrder()
	if err != nil {
		fmt.Println("ошибка при создании хранилища:", err)
		return
	}
	defer storOrder.Close()
	serviceOrder := service.NewOrder(&storOrder)

	id, err := strconv.Atoi(list[1])
	if err != nil {
		fmt.Println("неверно введено id", err)
		return
	}
	if id == 0 {
		log.Println("не указано id заказа")
		return
	}
	status := list[2]
	if status == "" {
		log.Println("не указано статус заказа")
		return
	}
	if parseStatus(status) == -1 {
		log.Println("статус заказа указан некоректно")
		return
	}
	err = serviceOrder.StatusOrder(model.OrderStatus{ID: id, Status: status}, chOrder)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("статус изменен")
}

func OrderSearch(list []string) {
	if len(list) != 3 {
		fmt.Println("Неправильно введена команда")
		return
	}
	muOrder.Lock()
	defer muOrder.Unlock()
	defer wg.Done()
	storOrder, err := storage.NewOrder()
	if err != nil {
		fmt.Println("ошибка при создании хранилища:", err)
		return
	}
	defer storOrder.Close()
	serviceOrder := service.NewOrder(&storOrder)

	fullname := list[1]
	if fullname == "" {
		log.Println("не указано ФИО")
		return
	}
	code := list[2]
	if code == "" {
		log.Println("не указан код заказа")
		return
	}

	err = serviceOrder.SearchOrder(model.OrderInput{FullName: fullname, OrderCode: code})
	if err != nil {
		fmt.Println(err)
		return
	}
}
