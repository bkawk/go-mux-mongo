package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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


func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", postUsers).Methods("POST")
	router.HandleFunc("/users", putUsers).Methods("PUT")
	router.HandleFunc("/users", deleteUsers).Methods("DELETE")
	router.HandleFunc("/users/{id}", findUsers).Methods("GET")

	return router
}