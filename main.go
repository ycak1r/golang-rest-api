package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Member struct (Model)
type Member struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book
var members []Member

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Add new book
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

// GET all clients
func getMembers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(members)
}

// Get single client
func getMember(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, memberItem := range members {
		if memberItem.ID == params["id"] {
			json.NewEncoder(w).Encode(memberItem)
			return
		}
	}
}

// Delete client
func deleteMember(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	for index, item := range members {
		if item.ID == param["id"] {
			members = append(members[:index], members[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(members)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// data @database books
	books = append(books, Book{ID: "1", Isbn: "234123", Title: "Book One", Author: &Author{Firstname: "Drogba", Lastname: "Bill"}})
	books = append(books, Book{ID: "2", Isbn: "343522", Title: "Book Two", Author: &Author{Firstname: "Fernando", Lastname: "Alpaçino"}})
	books = append(books, Book{ID: "3", Isbn: "123125", Title: "Book Three", Author: &Author{Firstname: "Jonathaan", Lastname: "Sneijder"}})

	// data @database clients
	members = append(members, Member{ID: "1", Name: "YCA", Type: "Tedarikçi"})
	members = append(members, Member{ID: "2", Name: "MEB", Type: "Müşteri"})
	members = append(members, Member{ID: "3", Name: "Eczane-Ekinoks", Type: "Her ikiside"})

	// Route handles & endpoints books
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	//  Route handles & endpoints members
	r.HandleFunc("/members", getMembers).Methods("GET")
	r.HandleFunc("/members/{id}", getMember).Methods("GET")
	r.HandleFunc("/members/{id}", deleteMember).Methods("DELETE")

	//Start server
	addr := ":8085"
	log.Println("Listen on ", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

// Request sample
// {
// 	"isbn":"123125",
// 	"title":"Book Three",
// 	"author":{"firstname":"Jonathaan","lastname":"Sneijder"}
// }

// Request client
// {
// 	"name":"YCA",
// 	"Type":"Tedarikçi",
// }
