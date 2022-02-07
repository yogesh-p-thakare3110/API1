package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting API server")
	//Create a new router
	r := mux.NewRouter()
	log.Println("creating routes")
	//Specify end points
	r.HandleFunc("/helloworld", PrintHelloWorld).Methods("GET")
	r.HandleFunc("/passed", printpassedtest).Methods("GET")
	http.Handle("/", r)

	//Start and listen to requests
	http.ListenAndServe(":8080", r)

}

func PrintHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "helloworld")
}

func printpassedtest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "PassedTest")
}
