package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// I combined both the TCP projects into one
func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Something went wrong")
		log.Fatalln(err.Error())
	}
	defer li.Close()
	fmt.Println("Line is closed")

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println("Uh oh")
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read (will read the connection, we want to return the URL)
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			m := strings.Fields(ln)
			fmt.Println("URL:::", m[1])
			// so this is a weird router because normally it would look at the URL and divide by route THEN by method, but whatever, this is practice
			if m[0] == "GET" {
				handleGet(conn, m[1])
			}
		}
		if ln == "" {
			break
		}
		i++
	}
}

// handles get requests
func handleGet(conn net.Conn, m string) {
	switch m {
	case "/abc":
		abc(conn)
	case "/anna":
		anna(conn)
	default:
		// write (will return the URL -- first assignment)
		respond(conn)
	}
}

// Some random functions to handle different requests
func abc(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>You reached ABC</strong></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
func anna(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><h1>You're a queen!</h1></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
