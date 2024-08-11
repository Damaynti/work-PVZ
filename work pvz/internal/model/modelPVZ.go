package model

type PVZ struct {
	ID                 int
	Title              string
	Address            string
	ContactInformation string
}

type PVZInput struct {
	Title              string
	Address            string
	ContactInformation string
}
