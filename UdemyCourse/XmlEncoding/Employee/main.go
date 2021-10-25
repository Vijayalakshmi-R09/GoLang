package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/getAllEmp", GetallEmploy)
	log.Fatal(http.ListenAndServe("localhost:8082", nil))
}

type Employee struct {
	Name string `json:"EmpName" xml:"Ename"`
	Id   string `json:"EmplID" xml:"EID"`
	Age  int
}

func GetallEmploy(rw http.ResponseWriter, r *http.Request) {
	emp := []Employee{
		{Name: "Vj", Id: "E1001", Age: 24},
		{Name: "Anu", Id: "E1002", Age: 28},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(rw).Encode(emp)
	} else {
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(emp)
	}
}
