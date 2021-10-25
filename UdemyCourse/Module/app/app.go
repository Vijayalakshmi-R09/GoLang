package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()
	router.HandleFunc("/getAllEmp", getallEmploy).Methods(http.MethodGet)
	//url segmentation and only numbric contraint
	router.HandleFunc("/getEmp/{Id:[0-9]+}", getEmp).Methods(http.MethodGet)
	router.HandleFunc("/getAllEmp", CreateEmpl).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe("localhost:8082", router)) //no nill bec we use new craeted multiplexer not default one
}

func getEmp(w http.ResponseWriter, re *http.Request) {
	vars := mux.Vars(re)
	fmt.Fprint(w, vars["Id"])
}

func CreateEmpl(w http.ResponseWriter, re *http.Request) {
	fmt.Fprint(w, "Post request received")
}
