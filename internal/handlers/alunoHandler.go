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

func (h *AlunoHandler) GetAlunos(w http.ResponseWriter, r *http.Request) {
	alunos, err := h.service.GetAllAlunos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(alunos)
}

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

func (h *AlunoHandler) CreateAluno(w http.ResponseWriter, r *http.Request) {
	var aluno models.Aluno
	json.NewDecoder(r.Body).Decode(&aluno)

	if err := h.service.CreateAluno(&aluno); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(aluno)
}

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

func (h *AlunoHandler) DeleteAluno(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := h.service.DeleteAluno(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
