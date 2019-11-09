package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./slice/sages.gohtml"))
}

func main() {
	sages := []string{"Martin Luther King", "The Oracle of Delphi", "Buddha", "Maggie Smith", "Captain Janeway", "Scully", "Marcus Aurelius"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)

	}
}
