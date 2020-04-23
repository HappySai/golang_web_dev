package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

//create function map which will have functions inside
var fm = template.FuncMap{
	"uc":       strings.ToUpper,
	"ft":       firstThree,
	"fdateMDY": monthDayYear,
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.html"))
}

func main() {

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
	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
		time      time.Time
	}{
		sages,
		cars,
		time.Now(),
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", data)
	if err != nil {
		log.Fatal(err)
	}

}
