package storage

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"

	"example.com/mymodule/internal/model"
)

const storageNamePVZ = "storagePVZ"

type StoragePVZ struct {
	storagePVZ *os.File
}

func (s *StoragePVZ) Close() error {
	return s.storagePVZ.Close()
}

func NewPVZ() (StoragePVZ, error) {
	filePvz, err := os.OpenFile(storageNamePVZ, os.O_CREATE, 0777)
	if err != nil {
		return StoragePVZ{}, err
	}
	return StoragePVZ{storagePVZ: filePvz}, nil
}

// создает запись о новом ПВЗ
func (service *StoragePVZ) CreatePVZ(input model.PVZInput, chPVZ chan bool) error {
	all, err := service.listAllPVZ()
	if err != nil {
		return err
	}

	newPVZ := PVZDTO{
		Title:              input.Title,
		ReceptionTime:      time.Now(),
		ID:                 len(all) + 1,
		IsDel:              false,
		Address:            input.Address,
		ContactInformation: input.ContactInformation,
	}

	all = append(all, newPVZ)
	err = writenBytes(all, chPVZ)
	if err != nil {
		return err
	}
	return nil
}

func writenBytes(PVZs []PVZDTO, chPVZ chan bool) error {
	rawBytes, err := json.Marshal(PVZs)
	if err != nil {
		return err
	}

	err = os.WriteFile(storageNamePVZ, rawBytes, 0777)
	chPVZ <- true
	if err != nil {
		return err
	}
	return nil
}

// удаляет запись о ПВЗ
func (service *StoragePVZ) DeletePVZ(id int, chPVZ chan bool) error {
	all, err := service.listAllPVZ()
	if err != nil {
		return err
	}
	if id > len(all) {
		return errors.New("ПВЗ с таким id не существует")
	}
	for index, order := range all {
		if order.ID == id {
			all[index].IsDel = true
		}
	}
	err = writenBytes(all, chPVZ)
	if err != nil {
		return err
	}
	return nil
}

// возвращает все существующие ПВЗ
func (service *StoragePVZ) ListPVZ() ([]model.PVZ, error) {
	all, err := service.listAllPVZ()
	if err != nil {
		return nil, err
	}

	onlyActive := make([]model.PVZ, 0, len(all))
	for _, do := range all {
		if !do.IsDel {
			onlyActive = append(onlyActive, model.PVZ{
				ID:                 do.ID,
				Title:              do.Title,
				Address:            do.Address,
				ContactInformation: do.ContactInformation,
			})
		}
	}

	return onlyActive, nil
}

func (service *StoragePVZ) listAllPVZ() ([]PVZDTO, error) {
	reader := bufio.NewReader(service.storagePVZ)
	rawBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var pvzs []PVZDTO
	if len(rawBytes) == 0 {
		return pvzs, nil
	}

	err = json.Unmarshal(rawBytes, &pvzs)
	if err != nil {
		return nil, err
	}
	return pvzs, nil
}
