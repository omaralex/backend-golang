package dtos

import "time"

type DrugResponse struct {
	ID          *uint32   `json:"id"`
	Name        string    `json:"name"`
	Approved    bool      `json:"approved"`
	MiniDose    int       `json:"min_dose"`
	MaxDose     int       `json:"max_dose"`
	AvailableAt time.Time `json:"available_at"`
}
