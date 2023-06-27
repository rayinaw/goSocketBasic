package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Unable to create socket listener!")
	}
	defer ln.Close()

	fmt.Println("Socket Server is listening on port 8080...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error when accepting the request")
			continue
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	d := make([]byte, 1024)
	_, err := conn.Read(d)
	if err != nil {
		fmt.Println("An error occured when reading the data")
	}

	fmt.Println(string(d))

	msg := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello, world!"
	conn.Write([]byte(msg))
}
