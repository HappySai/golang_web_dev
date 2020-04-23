package main

type gator int

/*
type person struct {
	fname string
	lname string
}
type personUpd struct {
	fname   string
	lname   string
	favFood []string
}

func (g personUpd) walk() string {
	return fmt.Sprintln(g.fname, g.lname, "is walking")
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func (y truck) transportationDevice() string {
	return fmt.Sprintln("I am truck with", y.doors, "doors and I am durable")
}
func (z sedan) transportationDevice() string {
	return fmt.Sprintln("I am a sedan with", z.doors, "doors and I am luxurious")
}

type transportation interface {
	transportationDevice() string
}

func report(g transportation) {
	fmt.Println(g.transportationDevice())
}

func main() {
	s := []int{10, 20, 30, 40}

	for index, v := range s {
		fmt.Println(index, v)
	}

	for i := range s {
		fmt.Println(i)
	}

	//map: composite literal (string and int)
	m := map[string]int{
		"Sai Prasad": 50,
		"Sai Rohit":  16,
	}

	for k := range m {
		fmt.Println(k)
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for _, v := range m {
		fmt.Println(v)
	}

	//Struct Person
	p := person{
		"Sai",
		"Prasad",
	}
	fmt.Println(p)
	fmt.Println(p.lname)
	fmt.Println(p.fname)

	p1 := personUpd{
		"Sai",
		"Prasad",
		[]string{"carrors", "orange", "apple", "peaches"},
	}

	fmt.Println(p1.fname)
	fmt.Println(p1.lname)
	fmt.Println(p1.favFood)

	for i, v := range p1.favFood {
		fmt.Println(i, v)
	}
	w := p1.walk()
	fmt.Println(w)

	t1 := truck{
		vehicle{
			2, "red",
		},
		true,
	}
	fmt.Println(t1)
	fmt.Println(t1.color)
	fmt.Println(t1.transportationDevice())
	s1 := sedan{
		vehicle{
			4,
			"white",
		},
		false,
	}
	fmt.Println(s1)
	fmt.Println(s1.luxury)
	fmt.Println(s1.transportationDevice())

	report(s1)
	report(t1)

	var g1 gator
	g1 = 100
	fmt.Println("The value of g1", g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = int(g1)
	fmt.Println(x)
}
*/
