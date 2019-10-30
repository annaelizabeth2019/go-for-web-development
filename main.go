package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func main() {

	p1 := person{
		"Anna",
		"Peterson",
	}
	fmt.Println(p1)

	sa1 := secretAgent{
		person{
			"Femme",
			"Nikita",
		},
		true,
	}

	fmt.Println(sa1)
	saySomething(p1)
	saySomething(sa1)

}

func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println(p.fname, `says, "Good morning, Anna."`)
}
func (sa secretAgent) speak() {
	fmt.Println(sa.fname, `says, "today is a good day to die"`)
}
