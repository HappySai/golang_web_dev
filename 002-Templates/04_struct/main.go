package main

import (
	"log"
	"os"
	"text/template"
)

type sages struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	// tpl = template.Must(template.ParseFiles("index.html"))
	tpl = template.Must(template.ParseFiles("indexUpd.html"))
}
func main() {

	//Create object of Struct sages
	buddha := sages{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
