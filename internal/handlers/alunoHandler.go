package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FelipeAJdev/dev-cloud-challenge/internal/models"
	"github.com/FelipeAJdev/dev-cloud-challenge/internal/services"

	"github.com/gorilla/mux"
	logrus "github.com/sirupsen/logrus"
)

type AlunoHandler struct {
	service services.AlunoService
	logger  *logrus.Logger
}

func NewAlunoHandler(service services.AlunoService, logger *logrus.Logger) *AlunoHandler {
	return &AlunoHandler{service, logger}
}

// function to send standardized JSON responses
func (h *AlunoHandler) sendResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			h.logger.WithError(err).Error("Failed to encode response data")
		}
	}
}

// helper function to send JSON error responses
func (h *AlunoHandler) sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	h.sendResponse(w, statusCode, models.ErrorResponse{Message: message, Code: statusCode})
}

// GetAlunos retorna todos os alunos cadastrados
// @Summary Retorna a lista de alunos
// @Description Obtém a lista de TODOS os alunos cadastrados no sistema
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Aluno
// @Failure 500 {object} models.ErrorResponse "Erro interno no servidor"
// @Router /alunos [get]
func (h *AlunoHandler) GetAlunos(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Received request to get all students")
	alunos, err := h.service.GetAllAlunos()
	if err != nil {
		h.logger.WithError(err).Error("Failed to get all students")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Erro ao obter alunos")
		return
	}

	h.logger.Info("Successfully retrieved all students")
	h.sendResponse(w, http.StatusOK, alunos)
}

// GetAluno retorna um aluno específico
// @Summary Retorna um aluno pelo ID
// @Description Obtém os dados de um aluno específico pelo ID
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Aluno"
// @Success 200 {object} models.Aluno "Dados do Aluno"
// @Failure 400 {object} models.ErrorResponse "ID inválido"
// @Failure 404 {object} models.ErrorResponse "Aluno não encontrado"
// @Failure 500 {object} models.ErrorResponse "Erro interno no servidor"
// @Router /alunos/{id} [get]
func (h *AlunoHandler) GetAluno(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.WithField("id", vars["id"]).Error("Invalid ID format")
		h.sendErrorResponse(w, http.StatusBadRequest, "ID inválido")
		return
	}

	h.logger.WithField("id", id).Info("Received request to get a student by ID")

	aluno, err := h.service.GetAlunoByID(id)
	if err != nil {
		h.logger.WithField("id", id).WithError(err).Error("Failed to get student by ID")
		h.sendErrorResponse(w, http.StatusNotFound, "Aluno não encontrado")
		return
	}

	h.logger.WithField("id", id).Info("Successfully retrieved student by ID")
	h.sendResponse(w, http.StatusOK, aluno)
}

// CreateAluno cria um novo aluno
// @Summary Cria um novo aluno
// @Description Adiciona um novo aluno ao sistema
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param aluno body models.Aluno true "Dados do Aluno"
// @Success 201 {object} models.Aluno
// @Failure 400 {object} models.ErrorResponse "Dados do aluno inválidos"
// @Failure 500 {object} models.ErrorResponse "Erro interno no servidor"
// @Router /alunos [post]
func (h *AlunoHandler) CreateAluno(w http.ResponseWriter, r *http.Request) {
	var aluno models.Aluno
	if err := json.NewDecoder(r.Body).Decode(&aluno); err != nil {
		h.logger.WithError(err).Error("Failed to decode student data")
		h.sendErrorResponse(w, http.StatusBadRequest, "Dados do aluno inválidos")
		return
	}

	h.logger.WithField("student", aluno).Info("Received request to create a new student")

	if err := h.service.CreateAluno(&aluno); err != nil {
		h.logger.WithError(err).Error("Failed to create a new student")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Erro ao criar aluno")
		return
	}

	h.logger.WithField("student", aluno).Info("Successfully created a new student")
	h.sendResponse(w, http.StatusCreated, aluno)
}

// UpdateAluno atualiza os dados de um aluno
// @Summary Atualiza os dados de um aluno
// @Description Atualiza as informações de um aluno específico pelo ID
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Aluno"
// @Param aluno body models.Aluno true "Dados do Aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} models.ErrorResponse "Dados inválidos ou ID inválido"
// @Failure 404 {object} models.ErrorResponse "Aluno não encontrado"
// @Failure 500 {object} models.ErrorResponse "Erro interno no servidor"
// @Router /alunos/{id} [put]
func (h *AlunoHandler) UpdateAluno(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.WithField("id", vars["id"]).Error("Invalid ID format")
		h.sendErrorResponse(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var aluno models.Aluno
	if err := json.NewDecoder(r.Body).Decode(&aluno); err != nil {
		h.logger.WithError(err).Error("Failed to decode student data for update")
		h.sendErrorResponse(w, http.StatusBadRequest, "Dados inválidos do aluno")
		return
	}
	aluno.ID = id

	h.logger.WithField("student", aluno).Info("Received request to update student")

	if err := h.service.UpdateAluno(&aluno); err != nil {
		h.logger.WithError(err).Error("Failed to update student")
		h.sendErrorResponse(w, http.StatusInternalServerError, "Erro ao atualizar aluno")
		return
	}

	h.logger.WithField("student", aluno).Info("Successfully updated student")
	h.sendResponse(w, http.StatusOK, aluno)
}

// DeleteAluno deleta um aluno
// @Summary Deleta um aluno pelo ID
// @Description Remove um aluno específico pelo ID
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Aluno"
// @Success 204 "No Content"
// @Failure 404 {object} models.ErrorResponse "Aluno não encontrado"
// @Failure 500 {object} models.ErrorResponse "Erro interno no servidor"
// @Router /alunos/{id} [delete]
func (h *AlunoHandler) DeleteAluno(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.logger.WithField("id", vars["id"]).Error("Invalid ID format")
		h.sendErrorResponse(w, http.StatusBadRequest, "ID inválido")
		return
	}

	h.logger.WithField("id", id).Info("Received request to delete student")

	if err := h.service.DeleteAluno(id); err != nil {
		h.logger.WithField("id", id).WithError(err).Error("Failed to delete student")
		h.sendErrorResponse(w, http.StatusNotFound, "Aluno não encontrado")
		return
	}

	h.logger.WithField("id", id).Info("Successfully deleted student")
	w.WriteHeader(http.StatusNoContent)
}
