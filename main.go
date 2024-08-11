package main

import (
	"net/http"
	"os"

	_ "github.com/FelipeAJdev/dev-cloud-challenge/docs" // Importa os documentos gerados pelo swagger
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
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title API de Gestão de Alunos
// @version 1.0
// @description Esta é a documentação da API de Gestão de Alunos.

// @contact.name Felipe Macedo
// @contact.email felipealexandrej@gmail.com

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	log := initLogger()

	// Carrega o .env apenas se não estiver em produção
	if os.Getenv("ENV") != "PRODUCTION" {
		if err := godotenv.Load(".env"); err != nil {
			log.WithFields(logrus.Fields{
				"error": err,
			}).Fatal("Erro ao carregar arquivo .env")
		}
	}

	database := pgstore.InitDB()
	defer func() {
		if err := database.Close(); err != nil {
			log.WithFields(logrus.Fields{
				"error": err,
			}).Error("Erro ao fechar a conexão com o banco de dados")
		} else {
			log.Info("Conexão com o banco de dados fechada com sucesso")
		}
	}()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set in the environment")
	}

	runMigrations(databaseURL, log)

	alunoRepository := repository.NewAlunoRepository(database)
	alunoService := services.NewAlunoService(alunoRepository)
	alunoHandler := handlers.NewAlunoHandler(alunoService, log)

	router := mux.NewRouter()

	router.HandleFunc("/alunos", alunoHandler.GetAlunos).Methods("GET")
	router.HandleFunc("/alunos", alunoHandler.CreateAluno).Methods("POST")
	router.HandleFunc("/alunos/{id}", alunoHandler.GetAluno).Methods("GET")
	router.HandleFunc("/alunos/{id}", alunoHandler.UpdateAluno).Methods("PUT")
	router.HandleFunc("/alunos/{id}", alunoHandler.DeleteAluno).Methods("DELETE")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default fallback
	}
	log.Info("Servidor rodando na porta " + port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Erro ao iniciar o servidor HTTP")
	}

}

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
