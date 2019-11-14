package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Make  string
	Model string
	Doors int
}

func init() {
	_, err := template.ParseFiles("sages.gohtml")
	if err != nil {
		fmt.Println(err)
	}
	t1 := template.New("sages.gohtml")
	tpl = template.Must(t1.Funcs(fm).ParseFiles("sages.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {

	b := sage{
		Name:  "Buddha",
		Motto: "The belief in no beliefs",
	}

	s := sage{
		Name:  "Dana Scully",
		Motto: "There must be a scientific explanation",
	}

	t := car{
		Make:  "Toyota",
		Model: "Camry",
		Doors: 4,
	}

	h := car{
		Make:  "Honda",
		Model: "Accord",
		Doors: 4,
	}

	sages := []sage{b, s}
	cars := []car{t, h}

	data := struct {
		Wisdom []sage
		Cars   []car
	}{
		sages,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
}
