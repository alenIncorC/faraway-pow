package tcpserver

import (
	"fmt"
	"net"

	"faraway-tcp-server/pkg/server/tcpserver/handler/pow"
)

func Run() {
	fmt.Println("Starting Word of Wisdom TCP server...")

	// listen for incoming connections
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	// accept incoming connections and handle them
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go pow.HandleConnection(conn)
	}
}
