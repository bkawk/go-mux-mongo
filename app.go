package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func findUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func postUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func putUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func deleteUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", postUsers).Methods("POST")
	router.HandleFunc("/users", putUsers).Methods("PUT")
	router.HandleFunc("/users", deleteUsers).Methods("DELETE")
	router.HandleFunc("/users/{id}", findUsers).Methods("GET")

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}

// go run app.go
