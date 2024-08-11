package main

import (
	"fmt"
	"log"
	"strconv"

	"example.com/mymodule/internal/model"
	"example.com/mymodule/internal/service"
	"example.com/mymodule/internal/storage"
)

func PVZList() {
	muPVZ.Lock()
	defer muPVZ.Unlock()
	defer wg.Done()
	storPVZ, err := storage.NewPVZ()
	if err != nil {
		fmt.Println("ошибка при создании хранилища:", err)
		return
	}
	defer storPVZ.Close()

	servicePVZ := service.NewPVZ(&storPVZ)
	list, err := servicePVZ.ListPVZ()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", list)
}

func PVZDelete(list []string) {
	if len(list) != 2 {
		fmt.Println("Неправильно введена команда")
		return
	}
	muPVZ.Lock()
	defer muPVZ.Unlock()
	defer wg.Done()
	storPVZ, err := storage.NewPVZ()
	if err != nil {
		fmt.Println("ошибка при создании хранилища:", err)
		return
	}
	defer storPVZ.Close()
	servicePVZ := service.NewPVZ(&storPVZ)

	id, err := strconv.Atoi(list[1])
	if err != nil {
		fmt.Println("неверно введено id", err)
		return
	}

	if id == 0 {
		log.Println("не указано id ПВЗ")
		return
	}
	err = servicePVZ.DeletePVZ(id, chPVZ)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ПВЗ удален")
}

func PVZCreate(list []string) {
	if len(list) != 4 {
		fmt.Println("Неправильно введена команда")
		return
	}
	muPVZ.Lock()
	defer muPVZ.Unlock()
	defer wg.Done()
	storPVZ, err := storage.NewPVZ()
	if err != nil {
		fmt.Println("ошибка при создании хранилища:", err)
		return
	}
	defer storPVZ.Close()
	servicePVZ := service.NewPVZ(&storPVZ)

	title := list[1]
	if title == "" {
		log.Println("не указано название ПВЗ")
	}
	address := list[2]
	if address == "" {
		log.Println("не указан адрес ПВЗ")
		return
	}
	contactInformation := list[3]
	if contactInformation == "" {
		log.Println("не указаны контактные данные ПВЗ")
		return
	}

	err = servicePVZ.CreatePVZ(model.PVZInput{Title: title, Address: address, ContactInformation: contactInformation}, chPVZ)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ПВЗ добавлен")
}
