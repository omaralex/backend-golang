package vaccination

import "time"

type Vaccination struct {
	ID     uint32
	Name   string
	DrugId int
	Dose   int
	Date   time.Time
}
