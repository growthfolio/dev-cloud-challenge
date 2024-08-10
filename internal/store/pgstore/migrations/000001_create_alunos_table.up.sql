CREATE TABLE IF NOT EXISTS alunos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    idade INT NOT NULL,
    nota_primeiro_semestre FLOAT NOT NULL,
    nota_segundo_semestre FLOAT NOT NULL,
    nome_professor VARCHAR(100) NOT NULL,
    numero_sala INT NOT NULL
);
