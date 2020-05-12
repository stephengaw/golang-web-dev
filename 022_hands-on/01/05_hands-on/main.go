package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

// define and initialise template
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type pet struct {
	Name string
	Animal string
	Age int
}

type pets []pet


func main() {

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}


func index(w http.ResponseWriter, r *http.Request) {

	p := pets{
		pet{"Ali", "cat", 3},
		pet{"Rover", "dog", 5},
		pet{"Shakespeare", "turtle", 45},
	}

	err := tpl.Execute(w, p)
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Its a dog!</h1>")
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Its me!</h1>")
}