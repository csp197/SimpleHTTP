package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

type FlagStruct struct {
	PathDirectory string
}

const (
	CRLF = "\r\n"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	log.Printf("Server Running...")

	// Initialize empty FlagStruct
	flags := FlagStruct{}

	// Check if `--directory` flag is passed during runtime
	// ./your_program.sh --directory <dir_name />
	// [0] => ./your_program.sh
	// [1] => --directory
	// [2] => <dir_name />
	if len(os.Args) == 3 && os.Args[1] == "--directory" {
		// Set FlagStruct's PathDirectory field to the passed directory argument
		flags.PathDirectory = os.Args[2]
		log.Printf(
			"Directory flag detected, using %s",
			flags.PathDirectory)
	}

	// Creating a TCP listener at port 4221
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Fatalln("Failed to bind to port 4221", err.Error())
	}
	// Ensure we close the listener after we're done
	defer listener.Close()

	for {
		// Accept the incoming connection from the listener at the binded port 4221
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln("Error accepting connection: ", err.Error())
		}

		// Define go routine for concurrency support
		go connectionHandler(connection, flags)
	}
}

func connectionHandler(connection net.Conn, flags FlagStruct) {
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

	// Tokenize the request payload by CRLF for easy parsing
	requestBreakdown := strings.Split(requestPayload, CRLF)

	// For debugging...
	// for idx, val := range requestBreakdown {
	// 	log.Printf("[%d] => %s", idx, val)
	// }

	// Request Breakdown
	// Status line -
	// GET                          // HTTP method 														// 0
	// /index.html                  // Request target 													// 0
	// HTTP/1.1                     // HTTP version														// 0
	// \r\n                         // CRLF that marks the end of the request line

	// Headers -
	// Host: localhost:4221\r\n     // Header that specifies the server's host and port					// 1
	// User-Agent: curl/7.64.1\r\n  // Header that describes the client's user agent					// 2
	// Accept: */*\r\n              // Header that specifies which media types the client can accept	// 3
	// \r\n                         // CRLF that marks the end of the headers

	// Define an empty compression scheme header to hold the string value of any passed compression schemes
	compressionSchemeHeader := ""

	// Compile a regexp struct to find the `Accept-Encoding` header
	encodingRegexpStruct, err := regexp.Compile(`(?m)Accept-Encoding: (.*)`)
	if err != nil {
		log.Fatalln(err)
	}

	// Extract request encoding compression scheme from request using regex
	encodingHeaderMatch := encodingRegexpStruct.FindStringSubmatch(requestPayload)

	// If the regex returns a positive number of matches and `gzip` is a substring of the header...
	// TODO: Remove hardcoded gzip to support more compression scheme headers
	if len(encodingHeaderMatch) > 0 && strings.Contains(encodingHeaderMatch[1], "gzip") {
		// Then, redefine the compression scheme header for the response with the extracted scheme
		compressionSchemeHeader = fmt.Sprintf("Content-Encoding: gzip" + CRLF)
		// log.Fatalf("%s", compressionSchemeHeader)
	}

	// For debugging...
	// fmt.Printf("[%s]", requestPayload)
	// for idx, val := range encodingHeaderMatch {
	// 	log.Printf("++++[%d] => %s", idx, val)
	// }

	// 200 Response Message as a byte array
	responseMessage := []byte("HTTP/1.1 200 OK" + CRLF + compressionSchemeHeader)

	// Tokenize status line by space for easy parsing
	statusLineBreakdown := strings.Split(requestBreakdown[0], " ")

	// Extract HTTP method from status line
	httpMethod := statusLineBreakdown[0]

	// Map containing valid HTTP methods for server
	validHttpMethods := map[string]bool{
		"GET":  true,
		"POST": true,
		"PUT":  false,
		"HEAD": false}

	// If the extracted HTTP method is not one of the valid HTTP methods...
	if !validHttpMethods[httpMethod] {
		// Then, print the following log statement and exit the server with 1 status code
		log.Fatalln(fmt.Sprintf("HTTP method %s is not currently supported", httpMethod))
	}

	// Compile optimized Regex struct using pattern to easily parse the request target
	// TODO: further improvements here?
	// regexPattern := "(?m)(/echo)[/a-z0-9]+"
	// regexpStruct, err := regexp.Compile(regexPattern)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Extract request target from status line
	requestTarget := statusLineBreakdown[1]

	// Tokenize the request target by `/` for easy parsing
	requestTargetBreakdown := strings.Split(requestTarget, "/")

	// Response Breakdown
	// Status line -
	// HTTP/1.1 200 OK
	// \r\n                          // CRLF that marks the end of the status line

	// Headers -
	// Content-Type: text/plain\r\n  // Header that specifies the format of the response body
	// Content-Length: 3\r\n         // Header that specifies the size of the response body, in bytes
	// \r\n                          // CRLF that marks the end of the headers

	// Response body -
	// abc                           // The string from the request

	// Check if the request target is an endpoint with no params...
	//		/
	//	 [0]/[1]
	if len(requestTargetBreakdown) == 2 {
		if requestTarget == "/" { // Check if the request target is the root of the server...
			// Then add the proper CRLF ending to the pre-existing 200 response
			responseMessage = []byte(string(responseMessage) + CRLF)
		} else if requestTargetBreakdown[1] == "user-agent" { // Check if the request target is the `user-agent` endpoint
			// Compile a regex struct to extract the user agent header from the incoming request
			userAgentRegexpStruct, err := regexp.Compile(`(?m)User-Agent: (.*)`)
			if err != nil {
				log.Fatalln(err)
			}
			// Extract client's user agent from request using regex
			userAgentMatch := userAgentRegexpStruct.FindStringSubmatch(requestPayload)

			// For debugging...
			// fmt.Printf("[%s]", requestPayload)
			// for idx, val := range userAgentMatch {
			// 	log.Printf("[%d] => %s", idx, val)
			// }
			// log.Printf("%d => %s", len(userAgent), userAgent)

			// Extract User-Agent value from the regex match
			// [0] => User-Agent: {val}
			// [1] => {val}
			userAgent := userAgentMatch[1]

			// Redefine the response message with the extracted user agent
			// The Content-Length should subtracted by 1 because strings are
			// 0-Indexed
			responseMessage = []byte(string(responseMessage) + fmt.Sprintf("Content-Type: text/plain%sContent-Length: %d%s%s%s", CRLF, len(userAgent)-1, userAgent, CRLF, CRLF))

		} else { // If the request target is not supported...
			// Then redefine the response message as a 404 response
			responseMessage = []byte("HTTP/1.1 404 Not Found" + CRLF + CRLF)
		}
	} else { // Otherwise, the request target either has params or has levels...
		if requestTargetBreakdown[1] == "echo" { // If the request target is "echo"...

			// Tokenize request target by `/` and access the string input at index 2
			// /echo/{str}
			// [0]/[1]/[2]
			strInput := requestTargetBreakdown[2]

			if strings.Contains(string(responseMessage), "gzip") {
				gzippedBufferString, err := gzipify(strInput)
				if err != nil {
					log.Fatalln(err)
				}
				responseMessage = []byte(string(responseMessage) + fmt.Sprintf("Content-Type: text/plain%sContent-Length: %d%s%s%s", CRLF, len(gzippedBufferString), gzippedBufferString, CRLF, CRLF))
			} else {
				// Redefine the response message by adding the passed string with headers
				responseMessage = []byte(string(responseMessage) + fmt.Sprintf("Content-Type: text/plain%sContent-Length: %d%s%s%s", CRLF, len(strInput), strInput, CRLF, CRLF))

			}

			// If the request target is "files" and the directory flag is passed and the FlagStruct is not empty
		} else if requestTargetBreakdown[1] == "files" && (FlagStruct{} != flags) {

			// Tokenize request target by `/` and access the string input at index 2
			// /files/{filename}
			// [0]/[1]/[2]
			fileName := requestTargetBreakdown[2]

			// Concatenate the passed pathDirectory string and the fileName to get the location of the file to be read
			absoluteFilePath := flags.PathDirectory + fileName

			if httpMethod == "GET" { // If the HTTP method is a GET request, then...
				// Read file in as a string
				fileString, err := os.ReadFile(absoluteFilePath)
				if err != nil { // If the file does not exist, then log the error and return a 404 response
					log.Println(err)
					// Redefine the response message as a 404 error
					responseMessage = []byte("HTTP/1.1 404 Not Found" + CRLF + CRLF)
				} else { // Otherwise, the file does exist and return a 200 response
					// Redefine the response message by adding the passed string with the apt headers
					responseMessage = []byte(
						string(responseMessage) + fmt.Sprintf("Content-Type: application/octet-stream%sContent-Length: %d%s%s%s", CRLF, len(fileString), fileString, CRLF, CRLF))
				}
			} else if httpMethod == "POST" { // If the HTTP method is a POST request, then...

				// Extract request body from request payload breakdown
				requestBodyBreakdown := requestBreakdown[len(requestBreakdown)-1]

				// For debugging...
				// for idx, val := range requestBreakdown {
				// 	log.Printf("%d => %s", idx, val)
				// }
				// log.Println(len(requestBreakdown))
				// log.Fatalln(requestBreakdown)

				// Write a file at `absoluteFilePath` with the request body as a byte array with 0644 permissions
				err := os.WriteFile(absoluteFilePath, []byte(requestBodyBreakdown), 0644)
				if err != nil {
					log.Fatalln(err)
				}

				// Redefine the response message by adding the passed string with the apt headers
				responseMessage = []byte("HTTP/1.1 201 Created" + CRLF + CRLF)
			}
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

// Function to encode the payload using gzip compression
func gzipify(data string) (string, error) {

	// Declare a bytes buffer object
	var buffer bytes.Buffer
	// Define a new gzip writer, using the buffer as storage
	gzipWriter := gzip.NewWriter(&buffer)

	// If the gzip writer can write the incoming data returns an error, then...
	if _, err := gzipWriter.Write([]byte(data)); err != nil {
		// Return an empty string and the error
		return "", err
	}
	// Otherwise close the gzipWriter and if there's an error there, then return it with an empty string
	if err := gzipWriter.Close(); err != nil {
		return "", err
	}
	// If nothing fails, then return the buffer object, containing the gzipped string
	return buffer.String(), nil
}
