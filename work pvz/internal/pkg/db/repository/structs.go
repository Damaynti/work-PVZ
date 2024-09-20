package repository

import (
	"errors"
	"time"
)

var ErrorObjectNotFound = errors.New("not found")


type PVZ struct {
	Title              string
	ReceptionTime      time.Time
	ID                 int64
	IsDel              bool
	Address            string
	ContactInformation string
}

type Orders struct {
	FullName      string
	ReceptionTime time.Time
	ID            int64
	IsDel         bool
	Status        string
	OrderCode     string
	PVZID         int64
}
