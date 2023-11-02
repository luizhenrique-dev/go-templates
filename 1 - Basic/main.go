package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{Nome: "Go", CargaHoraria: 40}
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Nome do Curso: {{.Nome}}\nCarga Horária: {{.CargaHoraria}}h\n")
	err := tmp.Execute(os.Stdout, curso)

	if err != nil {
		panic(err)
	}
}
