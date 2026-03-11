//go:build ignore
// full working example of a simple web server

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// How does this work...
// main function begins with a call to http.HandleFunc
// tells http pkg to handle all requests to the web root ("/") with handler
// it then calls listen and serve, port 8080 on any interface
// response writer assembles the HTTP server's response
// by writing to it, we send data to HTTP client
//  The trailing [1:] means "create a sub-slice of Path from the 1st character to the end." This drops the leading "/" from the path name

// http://localhost:8080/monkeys
// Hi there, I love monkeys!


