package main

import (
	"fmt"
	"os"
	"text/template"
)

//Reading all the files from folder
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("files/*"))
}

func main() {
	fmt.Println("This is Template Programming")
	/*tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		os.Exit(1)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("Error while creating file")
	}
	defer nf.Close()

	// err = tpl.Execute(os.Stdout, nil)
	err = tpl.Execute(nf, nil)
	if err != nil {
		os.Exit(1)
	}
	*/

	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		os.Exit(1)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "first.gohtml", nil)
	if err != nil {
		os.Exit(1)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		os.Exit(1)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	if err != nil {
		os.Exit(1)
	}

}
