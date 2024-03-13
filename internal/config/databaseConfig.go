package config

import (
	"database/sql"
	"fmt"
)


const (
	dbURL = "DB_URL" 
)

type DatabaseConfig interface {
	Address() *sql.DB 
}

type databaseConfig struct {
	dbUrl *sql.DB
}

func NewDatabaseConfig() (*databaseConfig, error) {
	dbUrl, err := sql.Open("sql", dbURL)
	if err != nil {
		return nil, fmt.Errorf("Can not connect to DB: ", err) 
	}

	return &databaseConfig{
		dbUrl:dbUrl,
	}, nil
}

func (cfg *databaseConfig) Address() *sql.DB {
	return cfg.dbUrl 
}
