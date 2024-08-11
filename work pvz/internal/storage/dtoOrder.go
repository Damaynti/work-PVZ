package storage

import "time"

type OrderDTO struct {
	FullName      string
	ReceptionTime time.Time
	ID            int
	IsDel         bool
	Status        string
	OrderCode     string
}
