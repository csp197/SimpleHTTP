package main

import (
	"log"

	// Uncomment this block to pass the first stage
	"net"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	log.Printf("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Printf("Failed to bind to port 4221")
		// os.Exit(1)
	}
	// Ensure we close the listener after we're done
	defer listener.Close()

	connection, err := listener.Accept()
	if err != nil {
		log.Printf("Error accepting connection: ", err.Error())
		// os.Exit(1)
	}
	// Ensure we close the connection after we're done
	defer connection.Close()

	buffer := make([]byte, 1024)
	numberOfBytes, err := connection.Read(buffer)
	if err != nil {
		log.Printf("Error writing to buffer: ", err.Error())
		// os.Exit(1)
	}
	log.Printf("received %d bytes", numberOfBytes)
	log.Printf("received the following data: %s", string(buffer[:numberOfBytes]))

	responseMessage := []byte("HTTP/1.1 200 OK\r\n\r\n")
	numberOfBytes, err = connection.Write(responseMessage)
	if err != nil {
		log.Printf("Error responding: ", err.Error())
		// os.Exit(1)
	}
	log.Printf("sent %d bytes", numberOfBytes)
	log.Printf("sent the following data: %s", string(buffer[:numberOfBytes]))
}
