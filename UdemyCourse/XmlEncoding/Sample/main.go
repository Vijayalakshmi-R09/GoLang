package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"Name" xml:"Name" `
	City    string `json:"CTY" xml:"CTY" `
	Zipcode string `json:"ZIP" xml:"ZIP" `
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
	if r.Header.Get("Content-Type") == "appliction/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(rw).Encode(customers)
	} else {
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(customers)
	}
}
