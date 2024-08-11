package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Command int

var wg sync.WaitGroup
var chOrder chan bool
var chPVZ chan bool

const (
	CreateOrder Command = iota
	DeleteOrder
	ListOrder
	StatusOrder
	SearchOrder
	CreatePVZ
	DeletePVZ
	ListPVZ
	Exit
)

func parseCommand(input string) Command {
	switch input {
	case "create_order":
		return CreateOrder
	case "delete_order":
		return DeleteOrder
	case "list_order":
		return ListOrder
	case "status_order":
		return StatusOrder
	case "search_order":
		return SearchOrder
	case "create_PVZ":
		return CreatePVZ
	case "delete_PVZ":
		return DeletePVZ
	case "list_PVZ":
		return ListPVZ
	case "exit":
		return Exit
	default:
		return -1
	}
}

func main() {
	wg.Add(1)
	chOrder = make(chan bool, 1)
	chPVZ = make(chan bool, 1)
	chOrder <- true
	chPVZ <- true
	Signal(chOrder, chPVZ)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		list := strings.Split(scanner.Text(), " ")
		command := parseCommand(list[0])
		wg.Add(1)
		switch command {
		case CreateOrder:
			<-chOrder
			go OrderCreate(list)

		case DeleteOrder:
			<-chOrder
			go OrderDelete(list)

		case ListOrder:
			go OrderList()

		case StatusOrder:
			<-chOrder
			go OrderStatus(list)

		case SearchOrder:
			go OrderSearch(list)

		case CreatePVZ:
			<-chPVZ
			go PVZCreate(list)

		case DeletePVZ:
			<-chPVZ
			go PVZDelete(list)

		case ListPVZ:
			go PVZList()

		case Exit:
			wg.Done()
			return
		default:
			wg.Done()
			fmt.Println("неизвестная команда")
		}
	}
	wg.Wait()
}
