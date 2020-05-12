package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"html/template"
)

type OpenClose struct {
	Open float64
	Close float64
}

type MarketChanges []OpenClose

var mc MarketChanges

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	f , err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}

	lines, err :=  csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	for idx, l := range lines {

		if idx != 0 {
			o, err := strconv.ParseFloat(l[1], 64)
			if err != nil {
				log.Fatalln(err)
			}
			c, err := strconv.ParseFloat(l[4], 64)
			if err != nil {
				log.Fatalln(err)
			}
			mc = append(mc, OpenClose{
				o,
				c,
			})
		}
	}

	err = t.Execute(os.Stdout, mc)
	if err != nil {
		log.Fatalln(err)
	}

}
