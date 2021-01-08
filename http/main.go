package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "<html><h3>hello</h3></html>")
	fmt.Println("HELLO")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			_, _ = fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Create file server handler
	// fs := http.FileServer( http.Dir( "/Users/Pawel/go/src/hello" ) )

	// start HTTP server with `fs` as the default handler
	// log.Fatal(http.ListenAndServe( ":8090", fs ))

	log.Fatal(http.ListenAndServe(":8090", nil))
}
