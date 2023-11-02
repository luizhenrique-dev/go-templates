package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	tmp := template.Must(template.New("template.html").ParseFiles([]string{"template.html"}...))
	err := tmp.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Python", 50},
		{"Java", 60},
		{"C#", 60},
		{"Ruby", 20},
		{"TypeScript", 30},
	})

	if err != nil {
		panic(err)
	}
}
