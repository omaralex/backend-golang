package datasource

import (
	"backend-kata/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func New(cfg config.PostgreSqlConfig) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Server, cfg.Port, cfg.User, cfg.Password, cfg.Database)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(3)
	db.SetConnMaxLifetime(1)

	if err = db.Ping(); err != nil {
		log.Error().Msg("Not possible to communicate with DB")
		panic(err)
	}

	return db
}
