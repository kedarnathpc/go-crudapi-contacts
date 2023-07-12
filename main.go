package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var contacts []Contact

type Contact struct {
	ID        string  `json:"id"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Number    *Number `json:"details"`
}

type Number struct {
	Number  string `json:"mob.no."`
	Address string `json:"address"`
}

func getContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range contacts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contact.ID = strconv.Itoa(rand.Intn(100))
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contact)
}

func updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range contacts {
		if item.ID == params["id"] {
			contacts = append(contacts[:index], contacts[index+1:]...)
			var contact Contact
			_ = json.NewDecoder(r.Body).Decode(&contact)
			contact.ID = params["id"]
			contacts = append(contacts, contact)
			json.NewEncoder(w).Encode(contact)
			return
		}
	}
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range contacts {
		if item.ID == params["id"] {
			contacts = append(contacts[:index], contacts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(contacts)
}

func main() {
	contacts = append(contacts, Contact{ID: "1", Firstname: "Ethan", Lastname: "Johnson", Number: &Number{Number: "+1 (555) 123-4567", Address: "123 Main Street, Anytown, USA"}})
	contacts = append(contacts, Contact{ID: "2", Firstname: "Olivia", Lastname: "Anderson", Number: &Number{Number: "+1 (555) 987-6543", Address: "456 Elm Avenue, Cityville, USA"}})
	contacts = append(contacts, Contact{ID: "3", Firstname: "Liam", Lastname: "Thompson", Number: &Number{Number: "+1 (555) 555-1234", Address: "789 Oak Drive, Countryside, USA"}})

	r := mux.NewRouter()

	r.HandleFunc("/contacts/all", getContacts).Methods("GET")
	r.HandleFunc("/contacts/{id}", getContact).Methods("GET")
	r.HandleFunc("/contacts/create", createContact).Methods("POST")
	r.HandleFunc("/contacts/update/{id}", updateContact).Methods("PUT")
	r.HandleFunc("/contacts/delete/{id}", deleteContact).Methods("DELETE")

	fmt.Printf("Starting server at post :8000\n")

	log.Fatal(http.ListenAndServe(":8000", r))
}
