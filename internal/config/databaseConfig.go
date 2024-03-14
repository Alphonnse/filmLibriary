package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
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
	log.Println("Connecting to database...")

	dbUrl := os.Getenv(dbURL)
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("Can not connect to DB: %s", err)
	}

	if err := conn.Ping(); err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &databaseConfig{
		dbUrl: conn,
	}, nil
}

func (cfg *databaseConfig) Address() *sql.DB {
	return cfg.dbUrl
}
