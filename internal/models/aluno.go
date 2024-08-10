package models

type Aluno struct {
	ID                   int     `json:"id"`
	Nome                 string  `json:"nome"`
	Idade                int     `json:"idade"`
	NotaPrimeiroSemestre float64 `json:"nota_primeiro_semestre"`
	NotaSegundoSemestre  float64 `json:"nota_segundo_semestre"`
	NomeProfessor        string  `json:"nome_professor"`
	NumeroSala           int     `json:"numero_sala"`
}
