package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	tmp := template.New("content.html")
	// Cria diretiva ToUpper para que possa ser usada no template
	tmp.Funcs(template.FuncMap{"ToUpper": ToUpper})
	tmp = template.Must(tmp.ParseFiles(templates...))

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
