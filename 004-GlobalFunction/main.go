package main

import (
	"log"
	"os"
	"text/template"
)

type user struct {
	Name  string
	Motto string
	Admin bool
}

type compare struct {
	Score1 int
	Score2 int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}
func main() {

	xs := []string{"one", "two", "three", "four"}
	values := []int{10, 20, 30, 40, 50}

	u1 := user{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: false,
	}
	u2 := user{
		Name:  "Gandhi",
		Motto: "Be the change",
		Admin: true,
	}
	u3 := user{
		Name:  "",
		Motto: "No body",
		Admin: true}

	users := []user{u1, u2, u3}
	co1 := compare{
		10,
		20,
	}
	co2 := compare{
		50,
		60,
	}
	compares := []compare{co1, co2}
	// compares := []compare{co1}

	data := struct {
		Words   []string
		Lname   string
		Fname   string
		Age     int
		Numbers []int
		Wisdom  []user
		NumComp []compare
	}{
		xs,
		"Sai Prasad",
		"Pusarla",
		50,
		values,
		users,
		compares,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
