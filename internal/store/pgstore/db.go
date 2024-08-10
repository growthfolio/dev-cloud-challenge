package pgstore

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {

	dbUser := os.Getenv("WSRS_DATABASE_USER")
	dbPassword := os.Getenv("WSRS_DATABASE_PASSWORD")
	dbName := os.Getenv("WSRS_DATABASE_NAME")
	dbHost := os.Getenv("WSRS_DATABASE_HOST")
	dbPort := os.Getenv("WSRS_DATABASE_PORT")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao pingar o banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso!")
	return db
}
