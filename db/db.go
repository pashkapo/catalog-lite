package db

import (
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/pashkapo/catalog-lite/core"
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
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("could not ping DB... %v", err)
	}

	return &Database{db}, nil
}
