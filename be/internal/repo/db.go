package repo

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresDB struct {
	DB *sql.DB
}

// <  SINGLETON PATTERN >

var once sync.Once
var instance *PostgresDB

func GetDatabaseInstance() *PostgresDB {
	once.Do(func() {
		instance = NewPostgresDB()
	})
	return instance
}

// </ SINGLETON PATTERN >

func NewPostgresDB() *PostgresDB {
	user := os.Getenv("NOME")
	password := os.Getenv("AGANDSKODE")
	port := os.Getenv("PRISTAV")
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	if user == "" || password == "" || port == "" {
		panic("NOME, AGANDSKODE, DB_HOST and PRISTAV environment variables must be set")
	}
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=pptx_db sslmode=disable", user, password, host, port)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(fmt.Sprintf("Failed to open DB: %v", err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to ping DB: %v", err))
	}

	return &PostgresDB{
		DB: db,
	}
}
