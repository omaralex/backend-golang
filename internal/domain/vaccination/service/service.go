package service

import (
	"backend-kata/internal/domain/vaccination"
	"backend-kata/internal/domain/vaccination/repository"
	"context"
)

type Repository interface {
	SaveVaccination(vaccination vaccination.Vaccination) (*uint32, error)
	UpdateVaccination(vaccinationId uint32, vaccination vaccination.Vaccination) error
	GetAllVaccinations(limit int, offset int) (*[]vaccination.Vaccination, error)
	DeleteVaccination(vaccinationId uint32) error
}
type VaccinationService struct {
	repository Repository
}

func NewVaccinationService(repository repository.VaccinationRepository) VaccinationService {
	return VaccinationService{repository}
}
func (s VaccinationService) Create(ctx context.Context, vaccination vaccination.Vaccination) (*uint32, error) {
	vaccinationId, errorSaveVaccination := s.repository.SaveVaccination(vaccination)
	if errorSaveVaccination != nil {
		return nil, errorSaveVaccination
	}

	return vaccinationId, nil
}
func (s VaccinationService) Update(ctx context.Context, vaccinationId uint32, vaccination vaccination.Vaccination) error {
	errorUpdateVaccination := s.repository.UpdateVaccination(vaccinationId, vaccination)
	if errorUpdateVaccination != nil {
		return errorUpdateVaccination
	}

	return nil
}
func (s VaccinationService) GetAll(ctx context.Context, limit int, offset int) (*[]vaccination.Vaccination, error) {
	vaccinations, errorGetAllVaccinations := s.repository.GetAllVaccinations(limit, offset)
	if errorGetAllVaccinations != nil {
		return nil, errorGetAllVaccinations
	}

	return vaccinations, nil
}

func (s VaccinationService) Delete(ctx context.Context, vaccinationId uint32) error {
	errorDeleteVaccination := s.repository.DeleteVaccination(vaccinationId)
	if errorDeleteVaccination != nil {
		return errorDeleteVaccination
	}

	return nil
}
