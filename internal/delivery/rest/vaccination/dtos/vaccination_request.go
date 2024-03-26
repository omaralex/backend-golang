package dtos

import (
	"backend-kata/internal/domain/vaccination"
	"github.com/google/uuid"
	"time"
)

type VaccinationRequest struct {
	Name   string    `json:"name"`
	DrugId int       `json:"drug_id"`
	Dose   int       `json:"dose"`
	Date   time.Time `json:"date"`
}

func (request VaccinationRequest) ToVaccination() vaccination.Vaccination {
	return vaccination.Vaccination{
		ID:     uuid.New().ID(),
		Name:   request.Name,
		DrugId: request.DrugId,
		Dose:   request.Dose,
		Date:   request.Date,
	}
}
