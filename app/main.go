package main

import (
	"log"
	"os"
)

// main is the entry point of the Go program.
//
// It initializes a SimpleServer struct with the host "0.0.0.0" and port 4221.
// It then checks if the command line arguments include the "--directory" flag and a directory path.
// If the flag and path are present, it sets the SimpleServer's dirPath field to the passed directory argument.
// If the flag is not present, it logs a message indicating that the flag has not been passed.
// Finally, it starts the simple HTTP server by calling the start method of the SimpleServer struct.
func main() {

	server := SimpleServer{
		host: "0.0.0.0",
		port: 4221,
	}

	// ./run.sh --directory <dir_name />
	// [0] => ./your_program.sh
	// [1] => --directory
	// [2] => <dir_name />
	log.Printf("Checking if `--directory` flag has been passed")
	if len(os.Args) == 3 && os.Args[1] == "--directory" {
		// Set SimpleServer's dirPath field to the passed directory argument
		server.dirPath = os.Args[2]
		log.Printf("Directory flag detected, using %s", server.dirPath)
	} else {
		log.Printf("`--directory` flag has not been passed")
	}

	log.Printf("Starting a simple HTTP server")
	server.start()

}
