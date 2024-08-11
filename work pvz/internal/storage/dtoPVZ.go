package storage

import "time"

type PVZDTO struct {
	Title              string
	ReceptionTime      time.Time
	ID                 int
	IsDel              bool
	Address            string
	ContactInformation string
}
