package file_components

import (
	"fmt"
	"net"
)

func createFileServer() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started, waiting for connections...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}
		go sendFile(conn)
	}
}

func sendFile(conn net.Conn) {
	defer conn.Close()
}
