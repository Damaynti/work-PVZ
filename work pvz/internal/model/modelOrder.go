package model

type Order struct {
	ID        int
	FullName  string
	Status    string
	OrderCode string
}

type OrderInput struct {
	FullName  string
	OrderCode string
}

type OrderStatus struct {
	ID     int
	Status string
}
