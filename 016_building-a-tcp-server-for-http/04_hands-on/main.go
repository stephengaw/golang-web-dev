package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	i := 0
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			fs := strings.Fields(ln)
			//fmt.Printf("The requested URL is: %s", fs[1])

			if len(fs) < 1 {
				continue
			}

			method := fs[0]
			//fmt.Fprintf(os.Stdout, "%s\n", method)

			switch method {
			case "GET":
				body := fs[1]
				get_response(conn, body)
			case "POST":
				post_response(conn)
				fmt.Fprintf(os.Stdout, "POST to URL: %s", fs[1])
			}


			break
		}
		i++
	}
	defer conn.Close()
}


func get_response(conn net.Conn, body string) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n") 			// status line
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body)) 	// header
	fmt.Fprint(conn, "Content-Type: text/plain\r\n") 		// header
	fmt.Fprint(conn, "\r\n") 							// blank line; CRLF; carriage-return line-feed
	fmt.Fprint(conn, body)
}

func post_response(conn net.Conn) {
	body := "POSTING!\r\n"

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n") 			// status line
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body)) 	// header
	fmt.Fprint(conn, "Content-Type: text/plain\r\n") 		// header
	fmt.Fprint(conn, "\r\n") 							// blank line; CRLF; carriage-return line-feed
	fmt.Fprint(conn, body)
}
