package main

import (
	"fmt"
	"log"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	log.Printf("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Fatalln("Failed to bind to port 4221")
	}
	// Ensure we close the listener after we're done
	defer listener.Close()

	connection, err := listener.Accept()
	if err != nil {
		log.Fatalln("Error accepting connection: ", err.Error())
	}
	// Ensure we close the connection after we're done
	defer connection.Close()

	buffer := make([]byte, 1024)
	numberOfBytes, err := connection.Read(buffer)
	if err != nil {
		log.Fatalln("Error writing to buffer: ", err.Error())
	}
	requestPayload := string(buffer[:numberOfBytes])
	log.Printf("received %d bytes", numberOfBytes)
	log.Printf("received the following data: %s", requestPayload)

	requestBreakdown := strings.Split(requestPayload, " ")

	// GET                          // HTTP method
	// /index.html                  // Request target
	// HTTP/1.1                     // HTTP version
	// \r\n                         // CRLF that marks the end of the request line

	// // Headers
	// Host: localhost:4221\r\n     // Header that specifies the server's host and port
	// User-Agent: curl/7.64.1\r\n  // Header that describes the client's user agent
	// Accept: */*\r\n              // Header that specifies which media types the client can accept
	// \r\n                         // CRLF that marks the end of the headers

	responseMessage := []byte("HTTP/1.1 200 OK\r\n\r\n")

	httpMethod := requestBreakdown[0]
	if httpMethod != "GET" {
		log.Fatalln(fmt.Sprintf("HTTP method %s is not currently supported", httpMethod))
	}

	requestTarget := requestBreakdown[1]
	if requestTarget != "/" {
		// log.Fatalln(fmt.Sprintf("HTTP target %s is not available", requestTarget))
		// respond with a 404
		responseMessage = []byte("HTTP/1.1 404 Not Found\r\n\r\n")
	}

	numberOfBytes, err = connection.Write(responseMessage)
	if err != nil {
		log.Fatalln("Error responding: ", err.Error())
	}
	log.Printf("sent %d bytes", numberOfBytes)
	// log.Printf("sent the following data: %s", string(buffer[:numberOfBytes]))

}
