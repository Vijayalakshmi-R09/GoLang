package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//defines the routes --->registering the needed routes with endpoint
	// //another way
	// http.HandleFunc("/greet", greet )
	http.HandleFunc("/greet", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello! welcome to Udemy course RSET API with GoLang...")
	})

	//starting the server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

// func greet(rw http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(rw, "Hello! welcome to Udemy course RSET API with GoLang...")
// }
