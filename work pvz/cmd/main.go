package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"example.com/mymodule/internal/model"
	"example.com/mymodule/internal/service"
	"example.com/mymodule/internal/storage"
)

func main() {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	fullname := createCmd.String("fullname", "", "fullname for order")
	orderCode:=createCmd.String("code","","code for order")

	delCmd := flag.NewFlagSet("del", flag.ExitOnError)
	orderId := delCmd.Int("id", 0, "id for order")

	statusCmd:=flag.NewFlagSet("status",flag.ExitOnError)
	id := statusCmd.Int("id", 0, "ID заказа")
	status:=statusCmd.String("status","","status for order")

	searchCmd:=flag.NewFlagSet("search",flag.ExitOnError)
	search:=searchCmd.String("fullname","","fullname for order")
	code:=searchCmd.String("code","","code for order")

	if len(os.Args) < 2 {
		fmt.Println("необходимо указать команду")
		return
	}

	command := os.Args[1]

	stor, err := storage.New()
	if err != nil {
		fmt.Println("ошибка при создании хранилища:", err)
		return
	}
	serv := service.New(&stor)

	switch command {
	case "create":
		err := createCmd.Parse(os.Args[2:])
		if err != nil {
			log.Println("ошибка парсинга аргументов:", err)
			return
		}
		if *orderCode == "" {
			log.Println("не указан код товара")
			return
		}
		if *fullname == "" {
			log.Println("не указано ФИО")
			return
		}
		err = serv.Create(model.OrderInput{FullName: *fullname,OrderCode: *orderCode})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("заказ добавлен")
	case "del":
		err := delCmd.Parse(os.Args[2:])
		if err != nil {
			log.Println("ошибка парсинга аргументов:", err)
			return
		}
		if *orderId == 0 {
			log.Println("не указано id заказа")
			return
		}
		err = serv.Del(*orderId)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("заказ удален")

	case "list":
		list, err := serv.List()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", list)

	case "status":
		err = statusCmd.Parse(os.Args[2:])
		if err != nil {
			log.Println("ошибка парсинга аргументов:", err)
			return
		}
		if *id == 0 {
			log.Println("не указано id заказа")
			return
		}
		if *status == "" {
			log.Println("не указано статус заказа")
			return
		}
		switch *status{
		case "выдан":  break
		case "возврат": break
		case "брак": break
		default: 
			log.Println("статус заказа указан некоректно") 
			return
		}
		err = serv.Status(model.OrderStatus{ID: *id,Status: *status})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("статус изменен")

	case "search":
		err := searchCmd.Parse(os.Args[2:])
		if err != nil {
			log.Println("ошибка парсинга аргументов:", err)
			return
		}
		if *search == "" {
			log.Println("не указано ФИО")
			return
		}
		if *code == "" {
			log.Println("не указан код заказа")
			return
		}
		
		err=serv.Search(model.OrderInput{FullName: *search, OrderCode: *code})
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Println("неизвестная команда")
	}
}
