package main

import (
	"os"
	"text/template"
)

func main() {
	indexHandler()
}

var tpl *template.Template

func init() {
	// tpl = template.Must(template.ParseFiles("./slice/sages.gohtml"))
}

// City tbd
type City struct {
	Name       string
	Country    string
	Population int64
}

func indexHandler() {

	// cities := make([]*City, 0)

	// // is there an easier way to populate the city struct
	// city := City{
	// 	Name:       "Austin",
	// 	Country:    "US",
	// 	Population: 200,
	// }
	// // This is my ugly solution to dealing with nil value from Firestore

	// log.Printf("**** => &city has %#v before appending to cities ****", &city)
	// cities = append(cities, &city)

	// fmt.Printf("**** => cities outside {} has %d", len(cities))
	// tpl.Execute(os.Stdout, cities)
	os.Stdout.Write([]byte("hello world"))
}
