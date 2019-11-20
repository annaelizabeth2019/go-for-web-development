package main

import "text/template"

import "os"

import "log"

import "fmt"

type person struct {
	Name string
	Age  int
}

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

func main() {
	p1 := person{
		Name: "James Bond",
		Age:  42,
	}

	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Programming in Go", "4"},
				course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
				course{"CSCI-140", "Mobile Apps Using Go", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-50", "Advanced Go", "5"},
				course{"CSCI-190", "Advanced Web Programming with Go", "5"},
				course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
			},
		},
	}

	data := struct {
		Year   year
		Person person
	}{
		y,
		p1,
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
}
