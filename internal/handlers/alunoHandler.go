package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FelipeAJdev/dev-cloud-challenge/internal/models"
	"github.com/FelipeAJdev/dev-cloud-challenge/internal/services"

	"github.com/gorilla/mux"
)

type AlunoHandler struct {
	service services.AlunoService
}

func NewAlunoHandler(service services.AlunoService) *AlunoHandler {
	return &AlunoHandler{service}
}

// GetAlunos retorna todos os alunos cadastrados
// @Summary Retorna a lista de alunos
// @Description Obtém a lista de TODOS os alunos cadastrados no sistema
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Aluno
// @Router /alunos [get]
func (h *AlunoHandler) GetAlunos(w http.ResponseWriter, r *http.Request) {
	alunos, err := h.service.GetAllAlunos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(alunos)
}

// GetAluno retorna um aluno específico
// @Summary Retorna um aluno pelo ID
// @Description Obtém os dados de um aluno específico pelo ID
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Aluno"
// @Success 200 {object} models.Aluno
// @Router /alunos/{id} [get]
func (h *AlunoHandler) GetAluno(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	aluno, err := h.service.GetAlunoByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(aluno)
}

// CreateAluno cria um novo aluno
// @Summary Cria um novo aluno
// @Description Adiciona um novo aluno ao sistema
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param aluno body models.Aluno true "Dados do Aluno"
// @Success 200 {object} models.Aluno
// @Router /alunos [post]
func (h *AlunoHandler) CreateAluno(w http.ResponseWriter, r *http.Request) {
	var aluno models.Aluno
	json.NewDecoder(r.Body).Decode(&aluno)

	if err := h.service.CreateAluno(&aluno); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(aluno)
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
// @Router /alunos/{id} [put]
func (h *AlunoHandler) UpdateAluno(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var aluno models.Aluno
	json.NewDecoder(r.Body).Decode(&aluno)
	aluno.ID = id

	if err := h.service.UpdateAluno(&aluno); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(aluno)
}

// DeleteAluno deleta um aluno
// @Summary Deleta um aluno pelo ID
// @Description Remove um aluno específico pelo ID
// @Tags Alunos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Aluno"
// @Success 204 "No Content"
// @Router /alunos/{id} [delete]
func (h *AlunoHandler) DeleteAluno(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := h.service.DeleteAluno(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
