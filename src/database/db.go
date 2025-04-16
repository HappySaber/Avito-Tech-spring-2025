package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

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
		log.Fatal("Переменная окружения DB_PORT не установлена")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT")) // Преобразуем строку в int
	if err != nil {
		log.Fatalf("Ошибка при преобразовании порта: %v", err)
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
	var err error
	dbConfig := BuildDBConfig()
	DB, err = sql.Open("postgres", DbURL(dbConfig))
	if err != nil {
		log.Fatalf("Ошибка при проверке подключения к базе данных: %v", err)
	}
	log.Println("Успешно подключено к базе данных!")
}
