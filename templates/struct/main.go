package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/struct/sages.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief in no beliefs",
	}
	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
}
