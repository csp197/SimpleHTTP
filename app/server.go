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
	log.Printf("Server Running...")

	// Uncomment this block to pass the first stage

	// Creating a TCP listener at port 4221
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Fatalln("Failed to bind to port 4221")
	}
	// Ensure we close the listener after we're done
	defer listener.Close()
	// Accept the incoming connection from the listener at the binded port 4221
	connection, err := listener.Accept()
	if err != nil {
		log.Fatalln("Error accepting connection: ", err.Error())
	}
	// Ensure we close the connection after we're done
	defer connection.Close()

	// Initialize byte arr of size 1024 as a buffer to store request/response data
	buffer := make([]byte, 1024)
	// Read data from incoming connection into the buffer and return the number of bytes up to which the buffer is written
	numberOfBytes, err := connection.Read(buffer)
	if err != nil {
		log.Fatalln("Error writing to buffer: ", err.Error())
	}
	// Typecast the written buffer from byte arr to string to form the request payload
	requestPayload := string(buffer[:numberOfBytes])

	// Log out the number of bytes written and the request payload
	log.Printf("received %d bytes", numberOfBytes)
	log.Printf("received the following data: %s", requestPayload)

	// Tokenize the request payload by space for easy parsing
	requestBreakdown := strings.Split(requestPayload, " ")

	// Request Breakdown
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

	// 200 Response Message as a byte array
	responseMessage := []byte("HTTP/1.1 200 OK\r\n")

	// Extract HTTP Method to ensure that it's a GET verb because of limited support of http server
	httpMethod := requestBreakdown[0]
	if httpMethod != "GET" {
		log.Fatalln(fmt.Sprintf("HTTP method %s is not currently supported", httpMethod))
	}

	// Compile optimized Regex struct using pattern to easily parse the request target
	// TODO: further improvements here?
	// regexPattern := "(?m)(/echo)[/a-z0-9]+"
	// regexpStruct, err := regexp.Compile(regexPattern)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Extract request target
	requestTarget := requestBreakdown[1]

	// Tokenize the request target by `/` for easy parsing
	requestTargetBreakdown := strings.Split(requestTarget, "/")

	// Response Breakdown
	// Status line
	// HTTP/1.1 200 OK
	// \r\n                          // CRLF that marks the end of the status line

	// Headers
	// Content-Type: text/plain\r\n  // Header that specifies the format of the response body
	// Content-Length: 3\r\n         // Header that specifies the size of the response body, in bytes
	// \r\n                          // CRLF that marks the end of the headers

	// Response body
	// abc                           // The string from the request

	// Check if the request target is an endpoint with no params...
	// 	  /
	// [0]/[1]
	if len(requestTargetBreakdown) == 2 {
		if requestTarget == "/" { // Check if the request target is the root of the server...
			// Then add the proper CRLF ending to the pre-existing 200 response
			responseMessage = []byte(string(responseMessage) + "\r\n")
		} else if requestTargetBreakdown[1] == "user-agent" { // Check if the request target is the `user-agent` endpoint
			// Extract client's user agent from request
			userAgent := requestBreakdown[5]
			// Redefine the response message with the extracted user agent
			responseMessage = []byte(
				string(responseMessage) + fmt.Sprintf(
					"Content-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
					len(userAgent),
					userAgent))

		} else { // If the request target is not supported...
			// Then redefine the response message as a 404 response
			responseMessage = []byte("HTTP/1.1 404 Not Found\r\n\r\n")
		}
	} else { // Otherwise, the request target either has params or has levels...
		if requestTargetBreakdown[1] == "echo" {

			// Tokenize request target by `/` and access the string input at index 2
			// /echo/{str}
			// [0]/[1]/[2]
			strInput := requestTargetBreakdown[2]

			// Redefine the response message by adding the passed string with headers
			responseMessage = []byte(
				string(responseMessage) + fmt.Sprintf(
					"Content-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
					len(strInput),
					strInput))
		}
	}

	// Write the response message as a byte array to the connection and get the number of bytes sent out
	numberOfBytes, err = connection.Write(responseMessage)
	if err != nil {
		log.Fatalln("Error responding: ", err.Error())
	}
	// Log out the number of bytes and stringify the response message
	log.Printf("sent %d bytes", numberOfBytes)
	log.Printf("sent the following data: %s", string(responseMessage))

}
