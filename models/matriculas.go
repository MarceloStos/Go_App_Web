package models

import (
	"Aplicacao_Web/db"
)

type Matricula struct {
	IdAluno       int
	IdDisciplina  int
	DataMatricula string
}

func BuscaMatricula() []Matricula {
	db := db.ConectaComBancoDeDados()

	selectTodosAsMatriculas, err := db.Query("select * from matricula")
	if err != nil {
		panic(err.Error())
	}

	matricula := Matricula{}
	matriculas := []Matricula{}

	for selectTodosAsMatriculas.Next() {
		var idAluno, idDisciplina int
		var dataMatricula string

		err := selectTodosAsMatriculas.Scan(&idAluno, &idDisciplina, &dataMatricula)

		if err != nil {
			panic(err.Error())
		}

		matricula.IdAluno = idAluno
		matricula.IdDisciplina = idDisciplina
		matricula.DataMatricula = dataMatricula

		matriculas = append(matriculas, matricula)
	}
	defer db.Close()
	return matriculas
}

func CriarMatricula(idAluno, idDisciplina int, dataMatricula string) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into matricula(aluno_id, disciplina_id, data_matricula) values($1, $2, $3)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(idAluno, idDisciplina, dataMatricula)
	defer db.Close()
}

func DeletarMatricula(idAluno, idDisciplina, dataMatricula string) {
	db := db.ConectaComBancoDeDados()

	deletarMatricula, err := db.Prepare("DELETE FROM matricula WHERE aluno_id = $1 AND disciplina_id = $2 AND data_matricula = $3")
	if err != nil {
		panic(err.Error())
	}

	deletarMatricula.Exec(idAluno, idDisciplina, dataMatricula)
	defer db.Close()
}

func EditarMatricula(idAluno, idDisciplina, dataMatricula string) Matricula {
	db := db.ConectaComBancoDeDados()

	MatriculaAlvo, err := db.Query("select * from matricula WHERE aluno_id = $1 AND disciplina_id = $2 AND data_matricula = $3", idAluno, idDisciplina, dataMatricula)

	if err != nil {
		panic(err.Error())
	}

	MatriculaParaAtualizar := Matricula{}

	for MatriculaAlvo.Next() {
		var idAluno, idDisciplina int
		var dataMatricula string

		err = MatriculaAlvo.Scan(&idAluno, &idDisciplina, &dataMatricula)

		if err != nil {
			panic(err.Error())
		}

		MatriculaParaAtualizar.IdAluno = idAluno
		MatriculaParaAtualizar.IdDisciplina = idDisciplina
		MatriculaParaAtualizar.DataMatricula = dataMatricula
	}
	defer db.Close()
	return MatriculaParaAtualizar
}

func AtualizarMatricula(idAluno, idDisciplina int, dataMatricula string, alunoIDAntigo int, disciplinaIDAntigo int) {
	db := db.ConectaComBancoDeDados()

	AtualizaMatricula, err := db.Prepare("UPDATE matricula SET aluno_id = $1, disciplina_id = $2, data_matricula = $3 WHERE aluno_id = $4 AND disciplina_id = $5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaMatricula.Exec(idAluno, idDisciplina, dataMatricula, alunoIDAntigo, disciplinaIDAntigo)
	defer db.Close()
}
