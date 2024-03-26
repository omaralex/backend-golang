package repository

import (
	"backend-kata/internal/domain"
	"backend-kata/internal/domain/auth"
	"database/sql"
	"errors"
	"github.com/rs/zerolog/log"
)

type UserRepository struct {
	sqlExecutor domain.SQLExecutor
}

func NewUserRepository(sqlExecutor domain.SQLExecutor) UserRepository {
	return UserRepository{sqlExecutor: sqlExecutor}
}

func (r UserRepository) SaveUser(user auth.User) error {
	sqlStatement := `INSERT INTO "user" (id, name, email, password) VALUES ($1, $2, $3, $4)`

	if _, err := r.sqlExecutor.Exec(sqlStatement, user.ID, user.Name, user.Email, user.Password); err != nil {
		log.Err(err).Msg("Error inserting into database")
		return err
	}

	log.Info().Msg("New record inserted successfully")

	return nil
}

func (r UserRepository) GetUserByEmail(email string) (*auth.User, error) {
	sqlStatement := `SELECT id, name, email, password FROM "user" WHERE email = $1`

	var user auth.User

	errorExecuteQuery := r.sqlExecutor.QueryRow(sqlStatement, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if errorExecuteQuery != nil {
		if errors.Is(errorExecuteQuery, sql.ErrNoRows) {
			log.Info().Msg("User not found")

			return nil, nil
		}

		log.Err(errorExecuteQuery).Msg("Error querying database")

		return nil, errorExecuteQuery
	}

	return &user, nil
}
