package postgres

import (
	"context"

	"example.com/mymodule/internal/model"
	"example.com/mymodule/internal/pkg/db"
)

type OrderRepo struct {
	db *db.Database
}

func NewOrder(database *db.Database) *OrderRepo {
	return &OrderRepo{db: database}
}

func (r *OrderRepo) Add(ctx context.Context,order *model.OrderInput) (int64, error) {
	
	var id int64

	err:=r.db.ExecQueryRow(ctx,`INSERT INTO orders(fullname,ordercode,status) VALUES ($1,$2,$3) RETURNING id;`,order.FullName,order.OrderCode,"заказ на складе").Scan(&id)
	return id,err
}


func (r *OrderRepo) GetAllOrder(ctx context.Context) ([]model.Order, error) {
    var orderList []model.Order
	raws, err := r.db.Query(ctx, "SELECT id,fullname,isdel,status,ordercode FROM orders")
	if err != nil {
        return nil, err
    }
    defer raws.Close()
	for raws.Next() {
		var order model.OrderRead
        if err := raws.Scan(&order.ID,&order.FullName,&order.IsDel,&order.Status,&order.OrderCode); err != nil {
            return nil, err
        }
		if !order.IsDel{
			orderList=append(orderList,model.Order{ID: order.ID,FullName: order.FullName,Status: order.Status,OrderCode: order.OrderCode})
		}
    }

    if err := raws.Err(); err != nil {
		return nil, err
    }
    return orderList, nil
}

func (r *OrderRepo) DeleteOrder(ctx context.Context, id int64) error {
	_,err:=r.db.Exec(ctx,`UPDATE orders SET isDel = $1 WHERE id = $2;`,true,id)
	return err
}

func(r *OrderRepo) StatusOrder(ctx context.Context,orderStatus model.OrderStatus)error{
	_,err:=r.db.Exec(ctx,`UPDATE orders SET status = $1 WHERE id = $2;`,orderStatus.Status,orderStatus.ID)
	return err
}

func(r *OrderRepo) SearchOrder(ctx context.Context,input model.OrderSerch)(int64,error){
	var id int64
	var isDel bool
	err:=r.db.ExecQueryRow(ctx,"SELECT id,isdel FROM orders WHERE fullname = $1 AND ordercode = $2",input.FullName,input.OrderCode).Scan(&id,&isDel)
	if isDel{id=0}
	return id,err
}
