package drug

import "time"

type Drug struct {
	ID          uint32
	Name        string
	Approved    bool
	MinDose     int
	MaxDose     int
	AvailableAt time.Time
}
