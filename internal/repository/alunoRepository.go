package repository

import (
	"database/sql"

	"github.com/FelipeAJdev/dev-cloud-challenge/internal/models"
)

type AlunoRepository interface {
	GetAll() ([]models.Aluno, error)
	GetByID(id int) (*models.Aluno, error)
	Create(aluno *models.Aluno) error
	Update(aluno *models.Aluno) error
	Delete(id int) error
}

type alunoRepository struct {
	db *sql.DB
}

func NewAlunoRepository(db *sql.DB) AlunoRepository {
	return &alunoRepository{db}
}

func (r *alunoRepository) GetAll() ([]models.Aluno, error) {
	rows, err := r.db.Query("SELECT * FROM alunos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alunos []models.Aluno
	for rows.Next() {
		var aluno models.Aluno
		if err := rows.Scan(&aluno.ID, &aluno.Nome, &aluno.Idade, &aluno.NotaPrimeiroSemestre, &aluno.NotaSegundoSemestre, &aluno.NomeProfessor, &aluno.NumeroSala); err != nil {
			return nil, err
		}
		alunos = append(alunos, aluno)
	}

	return alunos, nil
}

func (r *alunoRepository) GetByID(id int) (*models.Aluno, error) {
	var aluno models.Aluno
	err := r.db.QueryRow("SELECT * FROM alunos WHERE id = $1", id).Scan(&aluno.ID, &aluno.Nome, &aluno.Idade, &aluno.NotaPrimeiroSemestre, &aluno.NotaSegundoSemestre, &aluno.NomeProfessor, &aluno.NumeroSala)
	if err != nil {
		return nil, err
	}
	return &aluno, nil
}

func (r *alunoRepository) Create(aluno *models.Aluno) error {
	err := r.db.QueryRow("INSERT INTO alunos (nome, idade, nota_primeiro_semestre, nota_segundo_semestre, nome_professor, numero_sala) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		aluno.Nome, aluno.Idade, aluno.NotaPrimeiroSemestre, aluno.NotaSegundoSemestre, aluno.NomeProfessor, aluno.NumeroSala).Scan(&aluno.ID)
	return err
}

func (r *alunoRepository) Update(aluno *models.Aluno) error {
	_, err := r.db.Exec("UPDATE alunos SET nome = $1, idade = $2, nota_primeiro_semestre = $3, nota_segundo_semestre = $4, nome_professor = $5, numero_sala = $6 WHERE id = $7",
		aluno.Nome, aluno.Idade, aluno.NotaPrimeiroSemestre, aluno.NotaSegundoSemestre, aluno.NomeProfessor, aluno.NumeroSala, aluno.ID)
	return err
}

func (r *alunoRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM alunos WHERE id = $1", id)
	return err
}
