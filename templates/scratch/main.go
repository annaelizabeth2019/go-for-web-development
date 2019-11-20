package main

import (
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.New("apple").Parse("Here is my template"))
	tpl.ExecuteTemplate(os.Stdout, "apple", nil)
}
