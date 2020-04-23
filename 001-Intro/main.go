package main

import "fmt"

var y int
var z int = 100

//Data Structure
type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

//func (receiver) identifier (parameters)(returns){<code>}
func (p person) speak() {
	fmt.Println(p.fname, p.lname, `Hi there,  Good morning`)

}
func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `Say, "Shaken, not stirred." `)
}

func (p person) print() {
	fmt.Println(p.fname, p.lname)
}
func main() {
	fmt.Println("Inside main function")

	x := 7
	y = 10
	fmt.Println(x, y, z)

	z = 200
	fmt.Println(z)

	//slice datastructure
	xi := []int{10, 20, 30, 40, 50}
	fmt.Println(xi)

	//map literal - This is nothing but Key Value pair
	m := map[string]int{
		"Sai Prasad":   50,
		"Rajalakshmi":  45,
		"Sai Rohit":    15,
		"Sai Shobitha": 8,
	}
	fmt.Println("The value inside the map", m)

	p1 := person{fname: "Sai Prasad", lname: "Pusarla"}
	p2 := person{fname: "Sai Rohit", lname: "Pusarla"}
	p1.print()
	p2.print()
	p1.speak()
	p2.speak()

	saOne := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	saOne.speak()

	//Using the interface and polymorphism
	saySomething(p1)
	saySomething(p2)
	saySomething(saOne)
}
