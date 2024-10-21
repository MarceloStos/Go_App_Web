package models

import (
	"Aplicacao_Web/db"
)

type Disciplina struct {
	Id        int
	Nome      string
	Descricao string
}

func BuscaDisciplina() []Disciplina {
	db := db.ConectaComBancoDeDados()

	selectTodosAsDisciplinas, err := db.Query("select * from disciplina order by id asc")
	if err != nil {
		panic(err.Error())
	}

	disciplina := Disciplina{}
	disciplinas := []Disciplina{}

	for selectTodosAsDisciplinas.Next() {
		var id int
		var nome, descricao string

		err := selectTodosAsDisciplinas.Scan(&id, &nome, &descricao)

		if err != nil {
			panic(err.Error())
		}

		disciplina.Id = id
		disciplina.Nome = nome
		disciplina.Descricao = descricao

		disciplinas = append(disciplinas, disciplina)
	}
	defer db.Close()
	return disciplinas
}

func CriarDisciplina(nome, descricao string) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into disciplina(nome, descricao) values($1, $2)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao)
	defer db.Close()
}

func DeletarDisciplina(idDisciplina string) {
	db := db.ConectaComBancoDeDados()

	deletarDisciplina, err := db.Prepare("delete from disciplina where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarDisciplina.Exec(idDisciplina)
	defer db.Close()
}

func EditarDisciplina(idDisciplina string) Disciplina {
	db := db.ConectaComBancoDeDados()

	disciplinaAlvo, err := db.Query("select * from disciplina where id=$1", idDisciplina)

	if err != nil {
		panic(err.Error())
	}

	disciplinaParaAtualizar := Disciplina{}

	for disciplinaAlvo.Next() {
		var id int
		var nome, descricao string

		err = disciplinaAlvo.Scan(&id, &nome, &descricao)

		if err != nil {
			panic(err.Error())
		}

		disciplinaParaAtualizar.Id = id
		disciplinaParaAtualizar.Nome = nome
		disciplinaParaAtualizar.Descricao = descricao
	}
	defer db.Close()
	return disciplinaParaAtualizar
}

func AtualizarDisciplina(id int, nome, descricao string) {
	db := db.ConectaComBancoDeDados()

	AtualizaDisciplina, err := db.Prepare("update disciplina set nome=$1, descricao=$2 where id=$3")
	if err != nil {
		panic(err.Error())
	}
	AtualizaDisciplina.Exec(nome, descricao, id)
	defer db.Close()
}
