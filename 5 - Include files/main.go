package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	tmp := template.Must(template.New("content.html").ParseFiles(templates...))
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
