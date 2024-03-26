package dtos

import "time"

type VaccinationResponse struct {
	ID     *uint32   `json:"id"`
	Name   string    `json:"name"`
	DrugId int       `json:"drug_id"`
	Dose   int       `json:"dose"`
	Date   time.Time `json:"date"`
}
