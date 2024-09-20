package model


type PVZ struct {
	ID                 int64
	Title              string
	Address            string
	ContactInformation string
}

type PVZInput struct {
	Title              string
	Address            string
	ContactInformation string
}

type PVZRead struct {
	Title              string
	ID                 int64
	IsDel              bool
	Address            string
	ContactInformation string
}
