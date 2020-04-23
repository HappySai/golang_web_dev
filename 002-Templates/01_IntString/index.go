package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}
func main() {
	// err := tpl.ExecuteTemplate(os.Stdout, "index.html", 50)
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", `Release self-focus; embrace other-focus`)
	if err != nil {
		log.Fatalln(err)
	}
}
