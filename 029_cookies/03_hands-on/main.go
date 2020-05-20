package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {

	// look for cookie first
	c, err := r.Cookie("stephens-cookie")

	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "stephens-cookie",
			Value: "0",
		}
	}

	vc, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Fprintln(w, "We've seen you before! visit count =", vc)  // DOES NOT WORK IF ADD THIS BEFORE SetCookie

	vc++

	c.Value = strconv.Itoa(vc)

	http.SetCookie(w, c)  // must be called before any other use of w. Maybe because it adds a cookie HEADER, set before body?

	fmt.Fprintf(w, "Updated cookie to latest visit count of %s\n", c.Value)
	//io.WriteString(w, c.Value)
}
