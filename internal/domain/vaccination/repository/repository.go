package repository

import (
	"backend-kata/internal/domain"
	"backend-kata/internal/domain/vaccination"
	"github.com/rs/zerolog/log"
)

type VaccinationRepository struct {
	sqlExecutor domain.SQLExecutor
}

func NewVaccinationRepository(sqlExecutor domain.SQLExecutor) VaccinationRepository {
	return VaccinationRepository{sqlExecutor: sqlExecutor}
}

func (r VaccinationRepository) SaveVaccination(vaccination vaccination.Vaccination) (*uint32, error) {
	var insertedID *uint32
	sqlStatement := `INSERT INTO vaccination (
                  id, name, drugId, dose, date
                  ) VALUES ($1, $2, $3, $4, $5)`

	err := r.sqlExecutor.QueryRow(
		sqlStatement,
		vaccination.ID,
		vaccination.Name,
		vaccination.DrugId,
		vaccination.Dose,
		vaccination.Date,
	).Scan(&insertedID)

	if err != nil {
		log.Err(err).Msg("Error inserting into database")
		return nil, err
	}

	log.Info().Msg("New record inserted successfully")

	return insertedID, nil
}

func (r VaccinationRepository) UpdateVaccination(vaccinationId uint32, vaccination vaccination.Vaccination) error {
	sqlStatement := `UPDATE vaccination
        SET name = $1, drugId = $2, dose = $3, date = $4
        WHERE id = $6`

	_, err := r.sqlExecutor.Exec(
		sqlStatement,
		vaccination.Name,
		vaccination.DrugId,
		vaccination.Dose,
		vaccination.Date,
		vaccinationId,
	)

	if err != nil {
		log.Err(err).Msg("Error updating into database")
		return err
	}

	log.Info().Msg("Record updated successfully")

	return nil
}

func (r VaccinationRepository) DeleteVaccination(vaccinationId uint32) error {
	sqlStatement := `DELETE FROM vaccination WHERE id = $1`

	_, err := r.sqlExecutor.Exec(sqlStatement, vaccinationId)

	if err != nil {
		log.Err(err).Msg("Error deleting into database")
		return err
	}

	log.Info().Msg("Record deleted successfully")

	return nil
}

func (r VaccinationRepository) GetAllVaccinations(limit int, offset int) (*[]vaccination.Vaccination, error) {
	sqlStatement := `SELECT * FROM vaccination LIMIT $1 OFFSET $2`

	rows, errorQuery := r.sqlExecutor.Query(sqlStatement, limit, offset)

	if errorQuery != nil {
		log.Err(errorQuery).Msg("Error scanning into database")
		return nil, errorQuery
	}

	var vaccinations []vaccination.Vaccination
	for rows.Next() {
		var v vaccination.Vaccination
		if err := rows.Scan(&v.ID, &v.Name, &v.DrugId, &v.Dose, &v.Date); err != nil {
			log.Err(err).Msg("Error to scan the row")
		}

		vaccinations = append(vaccinations, v)
	}

	return &vaccinations, nil
}
