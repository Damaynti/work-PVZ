package service

import (
	"context"
	"database/sql"
	"errors"

	"example.com/mymodule/internal/model"
	"example.com/mymodule/internal/pkg/db/repository"
	"example.com/mymodule/internal/pkg/db/repository/postgres"
)
type Service struct {
    Repo *postgres.PVZRepo
}

func NewService(repo *postgres.PVZRepo) *Service {
    return &Service{Repo: repo}
}

// создание
func(s *Service) CreatePVZ(ctx context.Context,input model.PVZInput) (int64 ,error) {

	if len(input.Title) == 0 {
		return 0,errors.New("пустое название ПВЗ")
	}
	if len(input.Address) == 0 {
		return 0,errors.New("пустой адрес ПВЗ")
	}
	if len(input.ContactInformation) == 0 {
		return 0,errors.New("не заполнены контактные данные")
	}

	id, err := s.Repo.Add(ctx, &input)
	return id,err
}

// удаление
func (s *Service) DeletePVZ(ctx context.Context,id int64) error {

	if id == 0 {
		return errors.New("нулевой id")
	}
	err := s.Repo.DeletePVZ(ctx, int64(id))
	return err
}

// лист
func (s *Service) ListPVZ(ctx context.Context) ([]model.PVZ, error) {
	pvzs, err := s.Repo.GetAllPVZ(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrorObjectNotFound
		}
		return nil, err
	}
	return pvzs, nil
}
