package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templat/*.html"))
}
func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", 50)
	if err != nil {
		log.Fatalln(err)
	}
}
