package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	checkport := os.Getenv("DB_PORT")

	if checkport == "" {
		log.Fatal("Environment variable DB_PORT is not set")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error converting port: %v", err)
	}

	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
	)
}

func Init() {
	time.Sleep(5 * time.Second)
	var err error
	dbConfig := BuildDBConfig()
	DB, err = sql.Open("postgres", DbURL(dbConfig))
	if err != nil {
		log.Fatalf("Error checking database connection: %v", err)
	}
	if err = EnsureTablesExist(DB); err != nil {
		return
	}
	log.Println("Successfully connected to the database!")
}

func EnsureTablesExist(db *sql.DB) error {
	tables := []struct {
		name string
		sql  string
	}{
		{
			name: "uuid extension",
			sql:  `CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`,
		},
		{
			name: "users",
			sql: `CREATE TABLE IF NOT EXISTS users (
				id         UUID PRIMARY KEY             DEFAULT uuid_generate_v4(),
				email      VARCHAR(255) UNIQUE NOT NULL,
				password   VARCHAR(255)        NOT NULL,
				role       VARCHAR(255)        NOT NULL,
				created_at TIMESTAMPTZ         NOT NULL DEFAULT NOW()
			)`,
		},
		{
			name: "pvz",
			sql: `CREATE TABLE IF NOT EXISTS pvz (
				id         UUID PRIMARY KEY      DEFAULT uuid_generate_v4(),
				city       VARCHAR(255) NOT NULL,
				created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
			)`,
		},
		{
			name: "receptions",
			sql: `CREATE TABLE IF NOT EXISTS receptions (
				id         UUID PRIMARY KEY      DEFAULT uuid_generate_v4(),
				pvz_id     UUID      NOT NULL REFERENCES pvz (id) ON DELETE CASCADE,
				status     VARCHAR(255) NOT NULL,
				created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
			)`,
		},
		{
			name: "products",
			sql: `CREATE TABLE IF NOT EXISTS products (
				id           UUID PRIMARY KEY      DEFAULT uuid_generate_v4(),
				reception_id UUID      NOT NULL REFERENCES receptions (id) ON DELETE CASCADE,
				created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
				type         VARCHAR(255) NOT NULL
			)`,
		},
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	for _, table := range tables {
		if _, err = tx.Exec(table.sql); err != nil {
			return fmt.Errorf("failed to create %s: %w", table.name, err)
		}
		log.Printf("Table/extension %s ensured", table.name)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
