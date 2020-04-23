package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}
func main() {

	//Create object of Struct sages
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}
	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatered but with love alone is healed.",
	}
	jesus := sage{
		Name:  "Jesus",
		Motto: "Love All.",
	}
	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}
	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}

	f := car{
		Manufacturer: "Fiat",
		Model:        "Linea",
		Doors:        4,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	cars := []car{f, c}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
