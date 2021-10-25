package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Employee struct {
	Name string `json:"EmpName"`
	Id   string `json:"EmplID"`
	Age  int
}

func main() {
	http.HandleFunc("/getAllEmp", func(rw http.ResponseWriter, r *http.Request) {
		empl := []Employee{
			{Name: "Vj", Id: "E1001", Age: 24},
			{Name: "Anu", Id: "E1002", Age: 28},
		}
		rw.Header().Add("Content-type", "application/json")
		json.NewEncoder(rw).Encode(empl)
	})

	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}
