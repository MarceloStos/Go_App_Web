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

	selectTodosAsMatriculas, err := db.Query("select * from matricula order by data_matricula asc")
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

func CriarMatricula(idAluno, idDisciplina int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into matricula(aluno_id, disciplina_id, data_matricula) values($1, $2, $3)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(idAluno, idDisciplina)
	defer db.Close()
}

func DeletarMatricula(idMatricula string) {
	db := db.ConectaComBancoDeDados()

	deletarMatricula, err := db.Prepare("delete from matricula where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarMatricula.Exec(idMatricula)
	defer db.Close()
}

func EditarMatricula(idMatricula string) Matricula {
	db := db.ConectaComBancoDeDados()

	MatriculaAlvo, err := db.Query("select * from matricula where id=$1", idMatricula)

	if err != nil {
		panic(err.Error())
	}

	MatriculaParaAtualizar := Matricula{}

	for MatriculaAlvo.Next() {
		var id, idAluno, idDisciplina int
		var dataMatricula string

		err = MatriculaAlvo.Scan(&id, &idAluno, &idDisciplina, &dataMatricula)

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

func AtualizarMatricula(id, idAluno, idDisciplina int, dataMatricula string) {
	db := db.ConectaComBancoDeDados()

	AtualizaMatricula, err := db.Prepare("update matricula set idAluno=$2, idDisciplina=$3, data_matricula=$4 where id=$1")
	if err != nil {
		panic(err.Error())
	}
	AtualizaMatricula.Exec(id, idAluno, idDisciplina, dataMatricula)
	defer db.Close()
}
