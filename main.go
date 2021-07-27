package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Book Model
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"id"`
	Title  string  `json:"id"`
	Author *Author `json:"id"`
}

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Init Book var as a slice Book struct
var books []Book

// Get all books
func getBooks(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(books)
}

// Get one book
func getBook(responseWriter http.ResponseWriter, request *http.Request) {

}

// Create a book
func createBook(responseWriter http.ResponseWriter, request *http.Request) {

}

// Update a book
func updateBook(responseWriter http.ResponseWriter, request *http.Request) {

}

// Delete a Book
func deleteBook(responseWriter http.ResponseWriter, request *http.Request) {

}

func main() {

	// Init Router
	router := mux.NewRouter()

	// Mock Data - Implement database
	books = append(books, Book{ID: "1", Isbn: "23678", Title: "Algorithms",
		Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "78678", Title: "Data Structures",
		Author: &Author{FirstName: "Tom", LastName: "Kamau"}})

	// Router Handlers / Endpoints
	router.HandleFunc("api/books", getBooks).Methods("GET")
	router.HandleFunc("api/book/{id}", getBook).Methods("GET")
	router.HandleFunc("api/books", createBook).Methods("POST")
	router.HandleFunc("api/book/{id}", updateBook).Methods("PUT")
	router.HandleFunc("api/book/{id}", deleteBook).Methods("DELETE")

	//Run the Server
	log.Fatal(http.ListenAndServe(":8000", router))
}
