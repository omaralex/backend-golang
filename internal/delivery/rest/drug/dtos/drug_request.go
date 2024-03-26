package dtos

import (
	"backend-kata/internal/domain/drug"
	"github.com/google/uuid"
	"time"
)

type DrugRequest struct {
	Name        string    `json:"name"`
	Approved    bool      `json:"approved"`
	MiniDose    int       `json:"min_dose"`
	MaxDose     int       `json:"max_dose"`
	AvailableAt time.Time `json:"available_at"`
}

func (request DrugRequest) ToDrug() drug.Drug {
	return drug.Drug{
		ID:          uuid.New().ID(),
		Name:        request.Name,
		Approved:    request.Approved,
		MinDose:     request.MiniDose,
		MaxDose:     request.MaxDose,
		AvailableAt: request.AvailableAt,
	}
}
