package model


type Order struct {
	ID        int64
	FullName  string
	Status    string
	OrderCode string
}

type OrderInput struct {
	FullName  string
	OrderCode string
}

type OrderStatus struct {
	ID     int64
	Status string
}

type OrderRead struct {
	FullName string
	ID            int64
	IsDel         bool
	Status        string
	OrderCode     string
}

type OrderSerch struct{
	FullName string
	OrderCode string
}