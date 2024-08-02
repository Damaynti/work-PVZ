package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"example.com/mymodule/internal/model"
)

const storageName = "storage"

type Storage struct {
	storage *os.File
}

func New() (Storage,error){
	file, err := os.OpenFile(storageName,os.O_CREATE, 0777)
	if err != nil {
		return Storage{},err
	}
	return Storage{storage: file},nil
}

// создает заказ 
func (s *Storage) Create(input model.OrderInput) error {
	all,err:=s.listAll()
	if err!=nil{
		return err
	}

	newOrder:=OrderDTO{
	FullName:      input.FullName,
	ReceptionTime: time.Now(),
	ID:            len(all)+1,
	IsDel:         false,
	Status: "заказ на складе",
	OrderCode: input.OrderCode,
	}

	all=append(all,newOrder)
	err=writenBytes(all)
	if err!=nil{
		return err
	}
 	return nil
}

func writenBytes(orders []OrderDTO)error{
	rawBytes,err:=json.Marshal(orders)
	if err!=nil{
		return err
	}

	err=os.WriteFile(storageName,rawBytes,0777)
	if err!=nil{
		return err
	}
 	return nil
}

//меняет статус заказа по id
func (s *Storage) Status(orderStatus model.OrderStatus)error{
	all,err:=s.listAll()
	if err!=nil{
		return err
	}
	for index ,order:=range all{
		if order.ID==orderStatus.ID{
			all[index].Status=orderStatus.Status
		}
	}
	err=writenBytes(all)
	if err!=nil{
		return err
	}
	return nil
}


// удаляет заказ
func (s *Storage) Del(id int)error{
	all,err:=s.listAll()
	if err!=nil{
		return err
	}
	for index ,order:=range all{
		if order.ID==id{
			all[index].IsDel=true
		}
	}
	err=writenBytes(all)
	if err!=nil{
		return err
	}
	return nil
}

//поиск заказа по ФИО
func (s *Storage) Search(input model.OrderInput)error{
	all,err:=s.List()
	if err!=nil{
		return err
	}
	var existsOrder bool=false
	for index ,order:=range all{
		if (order.FullName==input.FullName && order.OrderCode==input.OrderCode){
			existsOrder=true
			fmt.Println(all[index].ID)
		}
	}
	if !existsOrder{
		fmt.Println("На это имя заказ не существует")
	}
	return nil
}

// возвращает все существующие заказы
func (s *Storage) List() ([]model.Order, error){
	all,err:=s.listAll()
	if err !=nil{
		return nil,err
	}

	onlyActive:=make([]model.Order,0,len(all))
	for _,do:=range all{
		if !do.IsDel{
			onlyActive=append(onlyActive,model.Order{ 
				ID:do.ID,
				FullName:do.FullName,
				Status:do.Status,
				OrderCode: do.OrderCode,
			})
		}	
	}

	return onlyActive,nil
}

func (s *Storage) listAll() ([]OrderDTO, error){
	reader := bufio.NewReader(s.storage)
	rawBytes, err :=io.ReadAll(reader)
	if err!= nil{
		return nil,err
	}

	var orders []OrderDTO
	if len(rawBytes)==0{
		return orders,nil
	}

	err =json.Unmarshal(rawBytes,&orders)
	if err !=nil{
		return nil,err
	}
	return orders,nil
}
