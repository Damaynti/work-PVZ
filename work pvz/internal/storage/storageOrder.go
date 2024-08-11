package storage

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"example.com/mymodule/internal/model"
)

const storageNameOrder = "storageOrder"

type StorageOrder struct {
	storageOrder *os.File
}

func (s *StorageOrder) Close() error {
	return s.storageOrder.Close()
}

func NewOrder() (StorageOrder, error) {
	fileOrder, err := os.OpenFile(storageNameOrder, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return StorageOrder{}, err
	}
	return StorageOrder{storageOrder: fileOrder}, nil
}

// создает заказ
func (service *StorageOrder) CreateOrder(input model.OrderInput, chOrder chan bool) error {
	all, err := service.listAllOrder()
	if err != nil {
		return err
	}

	newOrder := OrderDTO{
		FullName:      input.FullName,
		ReceptionTime: time.Now(),
		ID:            len(all) + 1,
		IsDel:         false,
		Status:        "заказ на складе",
		OrderCode:     input.OrderCode,
	}

	all = append(all, newOrder)
	err = writenBytesOrder(all, chOrder)
	if err != nil {
		return err
	}
	return nil
}

func writenBytesOrder(orders []OrderDTO, chOrder chan bool) error {
	rawBytes, err := json.Marshal(orders)
	if err != nil {
		return err
	}
	err = os.WriteFile(storageNameOrder, rawBytes, 0777)
	chOrder <- true
	if err != nil {
		return err
	}
	return nil
}

// меняет статус заказа по id
func (service *StorageOrder) StatusOrder(orderStatus model.OrderStatus, chOrder chan bool) error {
	all, err := service.listAllOrder()
	if err != nil {
		return err
	}
	if orderStatus.ID > len(all) {
		return errors.New("заказа с таким id не существует")
	}
	for index, order := range all {
		if order.ID == orderStatus.ID {
			all[index].Status = orderStatus.Status
		}
	}
	err = writenBytesOrder(all, chOrder)
	if err != nil {
		return err
	}
	return nil
}

// удаляет заказ
func (service *StorageOrder) DeleteOrder(id int, chOrder chan bool) error {
	all, err := service.listAllOrder()
	if err != nil {
		return err
	}
	if id > len(all) {
		return errors.New("заказа с таким id не существует")
	}
	for index, order := range all {
		if order.ID == id {
			all[index].IsDel = true
		}
	}
	err = writenBytesOrder(all, chOrder)
	if err != nil {
		return err
	}
	return nil
}

// поиск заказа по ФИО
func (service *StorageOrder) SearchOrder(input model.OrderInput) error {
	all, err := service.ListOrder()
	if err != nil {
		return err
	}
	var existsOrder bool = false
	for index, order := range all {
		if order.FullName == input.FullName && order.OrderCode == input.OrderCode {
			existsOrder = true
			fmt.Println(all[index].ID)
		}
	}
	if !existsOrder {
		fmt.Println("На это имя заказ не существует")
	}
	return nil
}

// возвращает все существующие заказы
func (service *StorageOrder) ListOrder() ([]model.Order, error) {
	all, err := service.listAllOrder()
	if err != nil {
		return nil, err
	}

	onlyActive := make([]model.Order, 0, len(all))
	for _, do := range all {
		if !do.IsDel {
			onlyActive = append(onlyActive, model.Order{
				ID:        do.ID,
				FullName:  do.FullName,
				Status:    do.Status,
				OrderCode: do.OrderCode,
			})
		}
	}

	return onlyActive, nil
}

func (service *StorageOrder) listAllOrder() ([]OrderDTO, error) {
	reader := bufio.NewReader(service.storageOrder)
	rawBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var orders []OrderDTO
	if len(rawBytes) == 0 {
		return orders, nil
	}

	err = json.Unmarshal(rawBytes, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
