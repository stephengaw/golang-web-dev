package main

import (
	"log"
	"os"
	"text/template")

type hotel struct {Name, Address, City, Zip, Region string}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	h1 := hotel{
		Name: "Marriot",
		Address: "1 Main St",
		City: "LA",
		Zip: "0123",
		Region: "Southern", // Central, Northern
	}

	h2 := hotel{
		Name: "Hilton",
		Address: "1 Washington Ave",
		City: "San Fran",
		Zip: "5678",
		Region: "Central", // Central, Northern
	}

	h3 := hotel{
		Name: "Travel Lodge",
		Address: "1 Laffayette",
		City: "San Diego",
		Zip: "9911",
		Region: "Northern", // Central, Northern
	}

	h := hotels{
		h1,
		h2,
		h3,
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}

}

