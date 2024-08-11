package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Signal(chOrder chan bool, chPVZ chan bool) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println("\nПолучен сигнал:", sig)
		exit()
		os.Exit(0)
	}()
}

func exit() {
	fmt.Println("Выполнение сохранения данных перед выходом")
	if <-chOrder && <-chPVZ {
		return
	}
}
