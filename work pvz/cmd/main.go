package main

import (
	"context"
	"log"
	"net/http"
	"sync"

	"example.com/mymodule/internal/pkg/db"
	"example.com/mymodule/internal/pkg/db/repository/postgres"
	"example.com/mymodule/internal/server"
	"example.com/mymodule/internal/service"
)

//type Command int

const port = ":8000"
const queryParamKey = "key"


var wg sync.WaitGroup

/*const (
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
}*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	database, err := db.NewOn(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer database.GetPool(ctx).Close()
	PVZRepo := service.NewService(postgres.NewPVZ(database))
	OrderRepo:=service.NewOrder(postgres.NewOrder(database))
	implemotation := server.Server{Service: PVZRepo,ServiceOrder: OrderRepo}
	http.Handle("/", server.CreateRouter(ctx,implemotation))
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

	

	/*wg.Add(1)
	signal.Signal()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		list := strings.Split(scanner.Text(), " ")
		command := parseCommand(list[0])
		wg.Add(1)
		switch command {
		case CreateOrder:
			//go comand.OrderCreate(ctx,list,&wg)

		case DeleteOrder:
			//go comand.OrderDelete(ctx,list,&wg)

		case ListOrder:
			//go comand.OrderList(ctx,&wg)

		case StatusOrder:
			//go comand.OrderStatus(ctx,list,&wg)

		case SearchOrder:
			//go comand.OrderSearch(ctx,list,&wg)

		case CreatePVZ:
			//go comand.PVZCreate(ctx,list,&wg)

		case DeletePVZ:
			//go comand.PVZDelete(ctx,list,&wg)

		case ListPVZ:
			//go comand.PVZList(ctx,&wg)

		case Exit:
			wg.Done()
			return
		default:
			wg.Done()
			fmt.Println("неизвестная команда")
		}
	}
	wg.Wait()*/
}
