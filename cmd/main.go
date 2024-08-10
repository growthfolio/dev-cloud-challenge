package main

import (
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
	logrus "github.com/sirupsen/logrus"
)

// @title API de Gestão de Alunos
// @version 1.0
// @description Esta é a documentação da API de Gestão de Alunos.

// @contact.name Felipe Macedo
// @contact.email felipealexandrej@gmail.com

// @host localhost:8080
// @BasePath /
// @schemes http

func initLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
	return log
}

func runMigrations(databaseURL string, log *logrus.Logger) {
	m, err := migrate.New(
		"file://./internal/store/pgstore/migrations",
		databaseURL)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
			"url":   databaseURL,
		}).Fatal("Falha ao criar a instância de migrate")
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Falha ao aplicar migrations")
	}

	log.Info("Migrations aplicadas com sucesso!")
}

func main() {
	log := initLogger()

	if err := godotenv.Load(); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Erro ao carregar arquivo .env")
	}

	database := pgstore.InitDB()
	defer func() {
		if err := database.Close(); err != nil {
			log.WithFields(logrus.Fields{
				"error": err,
			}).Error("Erro ao fechar a conexão com o banco de dados")
		}
	}()

	databaseURL := "postgres://postgres:123456789@localhost:5432/wsrs?sslmode=disable"

	runMigrations(databaseURL, log)

	alunoRepository := repository.NewAlunoRepository(database)
	alunoService := services.NewAlunoService(alunoRepository)
	alunoHandler := handlers.NewAlunoHandler(alunoService)

	router := mux.NewRouter()

	router.HandleFunc("/alunos", alunoHandler.GetAlunos).Methods("GET")
	router.HandleFunc("/alunos", alunoHandler.CreateAluno).Methods("POST")
	router.HandleFunc("/alunos/{id}", alunoHandler.GetAluno).Methods("GET")
	router.HandleFunc("/alunos/{id}", alunoHandler.UpdateAluno).Methods("PUT")
	router.HandleFunc("/alunos/{id}", alunoHandler.DeleteAluno).Methods("DELETE")

	log.Info("Servidor rodando na porta 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Erro ao iniciar o servidor HTTP")
	}
}
