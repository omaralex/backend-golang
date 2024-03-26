package service

import (
	"backend-kata/internal/domain/drug"
	"backend-kata/internal/domain/drug/repository"
	"context"
)

type Repository interface {
	SaveDrug(drug drug.Drug) (*uint32, error)
	UpdateDrug(drugId uint32, drug drug.Drug) error
	GetAllDrugs(limit int, offset int) (*[]drug.Drug, error)
	DeleteDrug(drugId uint32) error
}
type DrugService struct {
	repository Repository
}

func NewDrugService(repository repository.DrugRepository) DrugService {
	return DrugService{repository}
}
func (s DrugService) Create(ctx context.Context, drug drug.Drug) (*uint32, error) {
	id, errorSaveDrug := s.repository.SaveDrug(drug)
	if errorSaveDrug != nil {
		return nil, errorSaveDrug
	}

	return id, nil
}
func (s DrugService) Update(ctx context.Context, drugId uint32, drug drug.Drug) error {
	errorUpdateDrug := s.repository.UpdateDrug(drugId, drug)
	if errorUpdateDrug != nil {
		return errorUpdateDrug
	}

	return nil
}
func (s DrugService) GetAll(ctx context.Context, limit int, offset int) (*[]drug.Drug, error) {
	drugs, errorGetAllDrugs := s.repository.GetAllDrugs(limit, offset)
	if errorGetAllDrugs != nil {
		return nil, errorGetAllDrugs
	}

	return drugs, nil
}

func (s DrugService) Delete(ctx context.Context, drugId uint32) error {
	errorDeleteDrug := s.repository.DeleteDrug(drugId)
	if errorDeleteDrug != nil {
		return errorDeleteDrug
	}

	return nil
}
