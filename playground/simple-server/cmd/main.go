package main

import (
	"fmt"
	"log"
	"net/http"
)

// main starts a simple HTTP server on port 8081.
// It responds to all requests with a "Hello from the simple server!" message.
func main() {
	// Define the handler function for incoming HTTP requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set the response header to indicate the content type
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		// Write the response body
		fmt.Fprintln(w, "Hello from the simple server!")
		fmt.Fprintf(w, "You requested: %s %s\n", r.Method, r.URL.Path)
	})

	// Define the server address
	addr := ":8081"

	// Print a message to the console indicating the server is starting
	log.Printf("Server listening on http://localhost%s\n", addr)

	// Start the HTTP server and log any errors that occur
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
