package repository

import (
	"backend-kata/internal/domain"
	"backend-kata/internal/domain/drug"
	"github.com/rs/zerolog/log"
)

type DrugRepository struct {
	sqlExecutor domain.SQLExecutor
}

func NewDrugRepository(sqlExecutor domain.SQLExecutor) DrugRepository {
	return DrugRepository{sqlExecutor: sqlExecutor}
}

func (r DrugRepository) SaveDrug(drug drug.Drug) (*uint32, error) {
	var insertedID *uint32
	sqlStatement := `INSERT INTO drug (
                  id, name, approved, minDose, maxDose, availableAt
                  ) VALUES ($1, $2, $3, $4, $5, $6)`

	err := r.sqlExecutor.QueryRow(
		sqlStatement,
		drug.ID,
		drug.Name,
		drug.Approved,
		drug.MinDose,
		drug.MaxDose,
		drug.AvailableAt,
	).Scan(&insertedID)

	if err != nil {
		log.Err(err).Msg("Error inserting into database")
		return nil, err
	}

	log.Info().Msg("New record inserted successfully")

	return insertedID, nil
}

func (r DrugRepository) UpdateDrug(drugId uint32, drug drug.Drug) error {
	sqlStatement := `UPDATE drug
        SET name = $1, approved = $2, minDose = $3, maxDose = $4, availableAt = $5
        WHERE id = $6`

	_, err := r.sqlExecutor.Exec(
		sqlStatement,
		drug.Name,
		drug.Approved,
		drug.MinDose,
		drug.MaxDose,
		drug.AvailableAt,
		drugId,
	)

	if err != nil {
		log.Err(err).Msg("Error updating into database")
		return err
	}

	log.Info().Msg("Record updated successfully")

	return nil
}

func (r DrugRepository) DeleteDrug(drugId uint32) error {
	sqlStatement := `DELETE FROM drug WHERE id = $1`

	_, err := r.sqlExecutor.Exec(sqlStatement, drugId)

	if err != nil {
		log.Err(err).Msg("Error deleting into database")
		return err
	}

	log.Info().Msg("Record deleted successfully")

	return nil
}

func (r DrugRepository) GetAllDrugs(limit int, offset int) (*[]drug.Drug, error) {
	sqlStatement := `SELECT * FROM drug LIMIT $1 OFFSET $2`

	rows, errorQuery := r.sqlExecutor.Query(sqlStatement, limit, offset)

	if errorQuery != nil {
		log.Err(errorQuery).Msg("Error scanning into database")
		return nil, errorQuery
	}

	var drugs []drug.Drug
	for rows.Next() {
		var d drug.Drug
		if err := rows.Scan(&d.ID, &d.Name, &d.Approved, &d.MinDose, &d.MaxDose, &d.AvailableAt); err != nil {
			log.Err(err).Msg("Error to scan the row")
		}

		drugs = append(drugs, d)
	}

	return &drugs, nil
}
