package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	// tpl = template.Must(template.ParseFiles("index.html"))
	tpl = template.Must(template.ParseFiles("indexUpd.html"))
}
func main() {

	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad",
	}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
