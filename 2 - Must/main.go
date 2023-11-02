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
	tmp := template.Must(template.New("CursoTemplate").Parse("Nome do Curso: {{.Nome}}\nCarga Horária: {{.CargaHoraria}}h\n"))
	err := tmp.Execute(os.Stdout, curso)

	if err != nil {
		panic(err)
	}
}
