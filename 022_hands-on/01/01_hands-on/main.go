package main

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//w.Header("content-type: text/html")
	io.WriteString(w, `<a href="/dog">/dog</a><br><a href="/me">/me</a>`)
}

func dog(w http.ResponseWriter, r *http.Request) {
	//w.Header("content-type: text/html")
	io.WriteString(w, "<h1>Its a dog!</h1>")
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Its me!</h1>")
}


func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
