package pgstore

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Importa o driver do PostgreSQL
)

// InitDB inicializa e retorna uma conexão com o banco de dados PostgreSQL
func InitDB() *sql.DB {
	// Construir a string de conexão usando variáveis de ambiente
	dbUser := os.Getenv("WSRS_DATABASE_USER")
	dbPassword := os.Getenv("WSRS_DATABASE_PASSWORD")
	dbName := os.Getenv("WSRS_DATABASE_NAME")
	dbHost := os.Getenv("WSRS_DATABASE_HOST")
	dbPort := os.Getenv("WSRS_DATABASE_PORT")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Abrir a conexão com o banco de dados
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Verificar se a conexão foi bem-sucedida
	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao pingar o banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso!")
	return db
}
