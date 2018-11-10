package main

/*
executed command
go get github.com/gorilla/mux
*/

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Our main function
func main() {
	fmt.Println("My First API")

	router := mux.NewRouter()
	//implement endpoints for API
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

/**
 */
func GetPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

/**
 */
func GetPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

/**
 */
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

/**
 */
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
