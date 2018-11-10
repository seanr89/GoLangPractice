package main

/*
executed command
go get github.com/gorilla/mux
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////

/**
Request and return all people(JSON)
*/
func GetPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
	json.NewEncoder(w).Encode(people)
}

/**
 */
func GetPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
	//access params passed to API
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////

//Our main function
func main() {
	fmt.Println("My First API")

	router := mux.NewRouter()
	//Append dummy data to people array for api output
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Sean", Lastname: "Rafferty", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	//implement endpoints for API
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
