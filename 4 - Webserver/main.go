package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.New("template.html").ParseFiles([]string{"template.html"}...))
		err := tmp.Execute(w, Cursos{
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
	})

	http.ListenAndServe(":8082", nil)
}
