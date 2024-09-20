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


const port = ":8000"
const queryParamKey = "key"


var wg sync.WaitGroup



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
}
