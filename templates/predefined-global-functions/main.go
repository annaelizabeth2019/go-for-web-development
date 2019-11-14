package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type user struct {
	Name  string
	Motto string
	Admin bool
}

type scores struct {
	Score1 int
	Score2 int
}

func main() {

	a := user{
		Name:  "Anna",
		Motto: "Gradual progress gets results",
		Admin: true,
	}

	b := user{
		Name:  "Bart",
		Motto: "Eat my shorts!",
		Admin: false,
	}

	c := user{
		Name:  "Cat",
		Motto: "meow",
		Admin: true,
	}

	d := user{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}

	users := []user{a, b, c, d}

	xs := []string{"zero", "one", "two", "three", "four", "five"}

	g1 := scores{
		7,
		9,
	}

	data := struct {
		Words  []string
		Lname  string
		User   []user
		Scores scores
	}{
		xs,
		"Hsu",
		users,
		g1,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
}
