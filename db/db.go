package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/pashkapo/catalog-lite/config"
)

type Database struct {
	*sql.DB
}

func New(config *config.Config) (*Database, error) {
	databaseUrl := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=disable",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBName,
	)

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, errors.New(fmt.Sprintf("could not ping DB... %v", err))
	}

	return &Database{db}, nil
}
