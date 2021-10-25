package app

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	Name string `json:"EmpName" `
	Id   string `json:"EmplID" `
	Age  int
}

func getallEmploy(rw http.ResponseWriter, r *http.Request) {
	emp := []Employee{
		{Name: "Vj", Id: "E1001", Age: 24},
		{Name: "Anu", Id: "E1002", Age: 28},
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(emp)

}
