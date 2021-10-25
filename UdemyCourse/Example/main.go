package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/Names", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "\n1.viji")
		fmt.Fprint(rw, "\n2.Divi")
		fmt.Fprint(rw, "\n3.Vidi")
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
