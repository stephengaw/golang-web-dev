package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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

			body := fs[1]
			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n") 			// status line
			fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body)) 	// header
			fmt.Fprint(conn, "Content-Type: text/plain\r\n") 		// header
			fmt.Fprint(conn, "\r\n") 							// blank line; CRLF; carriage-return line-feed
			fmt.Fprint(conn, body)

			break
		}
		i++
	}
	defer conn.Close()
}
