package services

import (
	"github.com/FelipeAJdev/dev-cloud-challenge/internal/models"
	"github.com/FelipeAJdev/dev-cloud-challenge/internal/repository"
)

type AlunoService interface {
	GetAllAlunos() ([]models.Aluno, error)
	GetAlunoByID(id int) (*models.Aluno, error)
	CreateAluno(aluno *models.Aluno) error
	UpdateAluno(aluno *models.Aluno) error
	DeleteAluno(id int) error
}

type alunoService struct {
	repo repository.AlunoRepository
}

func NewAlunoService(repo repository.AlunoRepository) AlunoService {
	return &alunoService{repo}
}

func (s *alunoService) GetAllAlunos() ([]models.Aluno, error) {
	return s.repo.GetAll()
}

func (s *alunoService) GetAlunoByID(id int) (*models.Aluno, error) {
	return s.repo.GetByID(id)
}

func (s *alunoService) CreateAluno(aluno *models.Aluno) error {
	return s.repo.Create(aluno)
}

func (s *alunoService) UpdateAluno(aluno *models.Aluno) error {
	return s.repo.Update(aluno)
}

func (s *alunoService) DeleteAluno(id int) error {
	return s.repo.Delete(id)
}
