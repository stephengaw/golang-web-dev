package main

import (
	"os"
	"log"
	"text/template"
	)

type item struct {
	Service string
	Name string
	Price float64
}

type menu []item

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	m := menu{
		item{
			"Breakfast",
			"Eggs Bene",
			12.99,
		},
		item{
			"Lunch",
			"Soup",
			5.99,
		},
		item{
			"Lunch",
			"Club Sandwich",
			10.50,
		},
		item{
			"Dinner",
			"Steak",
			25.00,
		},
	}

	err := t.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
