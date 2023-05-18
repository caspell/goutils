package main

import (
	"fmt"
	"net"
)

func main() {
	// create a listener on port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	// accept incoming connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// handle incoming connection
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// read incoming data
	data := make([]byte, 1024)
	_, err := conn.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// print received data
	fmt.Printf("Received data: %s\n", string(data))

	// close the connection
	conn.Close()
}
