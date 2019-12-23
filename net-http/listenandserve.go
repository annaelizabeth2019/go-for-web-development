package main

import (
	"fmt"
	"net/http"
)

type cabbage int

// satisfies Handler interface
func (m cabbage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Anything you want")
}

func main() {
	var d cabbage
	http.ListenAndServe(":8080", d)
}
