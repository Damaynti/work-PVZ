package service

import (
	"errors"

	"example.com/mymodule/internal/model"
)

type storagePVZ interface {
	CreatePVZ(order model.PVZInput, chPVZ chan bool) error
	DeletePVZ(id int, chPVZ chan bool) error
	ListPVZ() ([]model.PVZ, error)
}

type ServicePVZ struct {
	servicePVZ storagePVZ
}

func NewPVZ(servicePVZ storagePVZ) ServicePVZ {
	return ServicePVZ{servicePVZ: servicePVZ}
}

// создание
func (service ServicePVZ) CreatePVZ(input model.PVZInput, chPVZ chan bool) error {

	if len(input.Title) == 0 {
		return errors.New("пустое название ПВЗ")
	}
	if len(input.Address) == 0 {
		return errors.New("пустой адрес ПВЗ")
	}
	if len(input.ContactInformation) == 0 {
		return errors.New("не заполнены контактные данные")
	}

	return service.servicePVZ.CreatePVZ(input, chPVZ)
}

// удаление
func (service ServicePVZ) DeletePVZ(id int, chPVZ chan bool) error {

	if id == 0 {
		return errors.New("нулевой id")
	}

	return service.servicePVZ.DeletePVZ(id, chPVZ)
}

// лист
func (service ServicePVZ) ListPVZ() ([]model.PVZ, error) {
	pvzs, err := service.servicePVZ.ListPVZ()
	if err != nil {
		return nil, err
	}
	return pvzs, nil
}
