package models

import (
	"Aplicacao_Web/db"
)

type Aluno struct {
	Id    int
	Nome  string
	Email string
}

func BuscaAluno() []Aluno {
	db := db.ConectaComBancoDeDados()

	selectTodosOsAlunos, err := db.Query("select * from aluno order by id asc")
	if err != nil {
		panic(err.Error())
	}

	aluno := Aluno{}
	alunos := []Aluno{}

	for selectTodosOsAlunos.Next() {
		var id int
		var nome, email string

		err := selectTodosOsAlunos.Scan(&id, &nome, &email)

		if err != nil {
			panic(err.Error())
		}

		aluno.Id = id
		aluno.Nome = nome
		aluno.Email = email

		alunos = append(alunos, aluno)
	}
	defer db.Close()
	return alunos
}

func CriarAluno(nome, email string) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into aluno(nome, email) values($1, $2)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, email)
	defer db.Close()
}

func DeletarAluno(idAluno string) {
	db := db.ConectaComBancoDeDados()

	deletarAluno, err := db.Prepare("delete from aluno where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarAluno.Exec(idAluno)
	defer db.Close()
}

func EditarAluno(idAluno string) Aluno {
	db := db.ConectaComBancoDeDados()

	AlunoAlvo, err := db.Query("select * from aluno where id=$1", idAluno)

	if err != nil {
		panic(err.Error())
	}

	AlunoParaAtualizar := Aluno{}

	for AlunoAlvo.Next() {
		var id int
		var nome, email string

		err = AlunoAlvo.Scan(&id, &nome, &email)

		if err != nil {
			panic(err.Error())
		}

		AlunoParaAtualizar.Id = id
		AlunoParaAtualizar.Nome = nome
		AlunoParaAtualizar.Email = email
	}
	defer db.Close()
	return AlunoParaAtualizar
}

func AtualizarAluno(id int, nome, email string) {
	db := db.ConectaComBancoDeDados()

	AtualizaAluno, err := db.Prepare("update aluno set nome=$2, email=$3 where id=$1")
	if err != nil {
		panic(err.Error())
	}
	AtualizaAluno.Exec(nome, email)
	defer db.Close()
}
