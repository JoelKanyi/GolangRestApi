package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book Model
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Model
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
	responseWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) // Get Parameters

	//Loop through the books and find the ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(responseWriter).Encode(item)
			return
		}
	}
	json.NewEncoder(responseWriter).Encode(&Book{})
}

// Create a book
func createBook(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID, cannot be used in production
	books = append(books, book)
	json.NewEncoder(responseWriter).Encode(book)
}

// Update a book - this reuses code for create and delete book
func updateBook(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(request.Body).Decode(&book)

			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(responseWriter).Encode(book)
			return
		}
		json.NewEncoder(responseWriter).Encode(books)
	}
}

// Delete a Book
func deleteBook(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
		json.NewEncoder(responseWriter).Encode(books)
	}
}

func main() {

	// Init Router
	router := mux.NewRouter()

	// Mock Data - Implement database
	books = append(books, Book{ID: "1", Isbn: "23678", Title: "Algorithms", Author: &Author{FirstName: "John", LastName: "Doe"}})

	books = append(books, Book{ID: "2", Isbn: "78678", Title: "Data Structures", Author: &Author{FirstName: "Tom", LastName: "Kamau"}})

	// Router Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//Run the Server
	log.Fatal(http.ListenAndServe(":8000", router))
}
