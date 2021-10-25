package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Customer struct {
	Name    string
	City    string
	Zipcode string
}

func main() {

	http.HandleFunc("/getAllCust", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}
func getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Samuel", City: "Chennai", Zipcode: "600096"},
		{Name: "Prasanth", City: "Chennai", Zipcode: "600096"},
		{Name: "Karthi", City: "Chennai", Zipcode: "600096"},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(customers)
}
