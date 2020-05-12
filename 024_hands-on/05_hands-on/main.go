package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("starting-files")))
	http.Handle("/pics/", http.FileServer(http.Dir("starting-files/public")))

	http.ListenAndServe(":8080", nil)
}
