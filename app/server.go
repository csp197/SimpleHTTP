package main

import (
	"fmt"
	"log"
	"regexp"
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

	// Status line
	// GET                          // HTTP method 														// 0
	// /index.html                  // Request target 													// 1
	// HTTP/1.1                     // HTTP version														// 2
	// \r\n                         // CRLF that marks the end of the request line						// 3

	// Headers
	// Host: localhost:4221\r\n     // Header that specifies the server's host and port					// 4
	// User-Agent: curl/7.64.1\r\n  // Header that describes the client's user agent					// 5
	// Accept: */*\r\n              // Header that specifies which media types the client can accept	// 6
	// \r\n                         // CRLF that marks the end of the headers							// 7

	responseMessage := []byte("HTTP/1.1 200 OK\r\n")

	httpMethod := requestBreakdown[0]
	if httpMethod != "GET" {
		log.Fatalln(fmt.Sprintf("HTTP method %s is not currently supported", httpMethod))
	}

	regexPattern := "(?m)(/echo)[/a-z0-9]+"
	regexpStruct, err := regexp.Compile(regexPattern)
	if err != nil {
		log.Fatalln(err)
	}

	requestTarget := requestBreakdown[1]

	respondWith404 := false

	strInput := ""

	if requestTarget != "/" {
		regexString := regexpStruct.FindString(requestTarget)
		if len(regexString) == 0 {
			respondWith404 = true
		} else {
			// TODO extract substring using regex
			// regexSubstringPattern := "(?m)(/echo)[/a-z0-9]+"
			// regexpStruct, err := regexp.Compile(regexPattern)
			strInput = strings.Split(regexString, "/")[2]
		}
	}
	if respondWith404 {
		responseMessage = []byte("HTTP/1.1 404 Not Found\r\n\r\n")
	} else if len(strInput) > 0 {
		responseMessage = []byte(string(responseMessage) + fmt.Sprintf("Content-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(strInput), strInput))
	}

	// Status line
	// HTTP/1.1 200 OK
	// \r\n                          // CRLF that marks the end of the status line

	// Headers
	// Content-Type: text/plain\r\n  // Header that specifies the format of the response body
	// Content-Length: 3\r\n         // Header that specifies the size of the response body, in bytes
	// \r\n                          // CRLF that marks the end of the headers

	// Response body
	// abc                           // The string from the request

	numberOfBytes, err = connection.Write(responseMessage)
	if err != nil {
		log.Fatalln("Error responding: ", err.Error())
	}
	log.Printf("sent %d bytes", numberOfBytes)
	log.Printf("sent the following data: %s", string(responseMessage))

}
