package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("amangos"))
	})

	r.Get("/{name}-{age}", func(w http.ResponseWriter, r *http.Request) {
		nameP := chi.URLParam(r, "name")
		ageP := chi.URLParam(r, "age")
		ageInt, err := strconv.Atoi(ageP)
		if err != nil {
			w.WriteHeader(422)
			w.Write([]byte(fmt.Sprintf("error parsing age: %s", ageP)))
		}

		fmt.Printf("name: %s\nage: %d", nameP, ageInt)

		/*
			response, err := json.Marshal(Person{nameP, ageInt})
			if err != nil {
				w.WriteHeader(422)
				w.Write([]byte("error making json or something"))
			}
		*/

		templ, err := template.New("personTemplate").Parse("<h1>Hi {{.Name}}, aged {{.Age}}</h1>")
		if err != nil {
			w.WriteHeader(422)
			w.Write([]byte("error making template"))
		}

		err = templ.Execute(w, Person{nameP, ageInt})
		if err != nil {
			w.WriteHeader(422)
			w.Write([]byte("error executing template"))
		}

		//w.Write()
		w.WriteHeader(200)
	})

	http.ListenAndServe("127.0.0.1:3000", r)
}
