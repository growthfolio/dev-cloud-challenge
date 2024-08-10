package main

import (
	"log"
	"net/http"

	"github.com/FelipeAJdev/dev-cloud-challenge/internal/handlers"
	"github.com/FelipeAJdev/dev-cloud-challenge/internal/repository"
	"github.com/FelipeAJdev/dev-cloud-challenge/internal/services"
	"github.com/FelipeAJdev/dev-cloud-challenge/internal/store/pgstore"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func runMigrations(databaseURL string) {
	m, err := migrate.New(
		"file://./internal/store/pgstore/migrations",
		databaseURL)
	if err != nil {
		log.Fatalf("Falha ao criar a instância de migrate: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Falha ao aplicar migrations: %v", err)
	}

	log.Println("Migrations aplicadas com sucesso!")
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}

	database := pgstore.InitDB()
	defer database.Close()

	databaseURL := "postgres://postgres:123456789@localhost:5432/wsrs?sslmode=disable"

	runMigrations(databaseURL)

	alunoRepository := repository.NewAlunoRepository(database)
	alunoService := services.NewAlunoService(alunoRepository)
	alunoHandler := handlers.NewAlunoHandler(alunoService)

	router := mux.NewRouter()

	router.HandleFunc("/alunos", alunoHandler.GetAlunos).Methods("GET")
	router.HandleFunc("/alunos", alunoHandler.CreateAluno).Methods("POST")
	router.HandleFunc("/alunos/{id}", alunoHandler.GetAluno).Methods("GET")
	router.HandleFunc("/alunos/{id}", alunoHandler.UpdateAluno).Methods("PUT")
	router.HandleFunc("/alunos/{id}", alunoHandler.DeleteAluno).Methods("DELETE")

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
