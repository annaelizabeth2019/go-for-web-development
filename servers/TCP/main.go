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

	// write (will return the URL)
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			m := strings.Fields(ln)
			fmt.Println("URL:::", m[1])
			if m[0] == "GET" {
				handleGet(conn, m)

			}
		}
		if ln == "" {
			break
		}
		i++
	}
}

func handleGet(conn net.Conn, m []string) {
	switch m[1] {
	case "/abc":
		abc(conn)
	case "/anna":
		anna(conn)
	default:
		// some error page
	}
}

func abc(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>You reached ABC</strong></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
func anna(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>You're a queen!</strong></body></html>`
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
