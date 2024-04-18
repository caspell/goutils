package main

import (
	"fmt"
	"net"
)

func main() {
	// connect to the server on port 8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// send data to the server
	data := []byte("Hello, world!")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
