package main

import (
	"encoding/json"
	"log" //Logging Errors
	"math/rand"
	"net/http" //Creation of ids
	"strconv"

	//String Converter

	"github.com/gorilla/mux" // Router
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`     // Fetch id property from json file
	Isbn   string  `json:"isbn"`   // Fetch Isbn property from json file
	Title  string  `json:"title"`  // Fetch Title property from json file
	Author *Author `json:"author"` // Fetch Author property from json file
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
// Slice is a variable length array
var books []Book

// Route Functions

// getBooks returns all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	// Set header type as application json
	w.Header().Set("Content-Type", "application/json")

	// Write books slice to response
	json.NewEncoder(w).Encode(books)
}

// getBooks returns single books
func getBook(w http.ResponseWriter, r *http.Request) {
	// Set Headers
	w.Header().Set("Content-Type", "application/json")

	// Get Parameters from request
	params := mux.Vars(r)

	//Loop through books and find id
	for _, item := range books {
		// Check if current ID is equal to requested id
		if item.ID == params["id"] {
			// Write single item to response
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// If nothing found, respond with empty Book object
	json.NewEncoder(w).Encode(&Book{})
}

// createBook creates single book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book) // Get book item from request
	// Create random id and convert to string
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Just an example - Not Production Safe

	// Append newly created book to collection
	books = append(books, book)

	// Add new book item to response
	json.NewEncoder(w).Encode(book)
}

// updateBook update single book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Search through books and when match is found
	//
	for index, item := range books {
		if item.ID == params["id"] {
			// Take slice up until (not including index)
			// and append remaining list
			// Effectively removes specific book
			books = append(books[:index], books[index+1:]...)

			// Now create new book with updates
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book) // Get book item from request

			// Create random id and convert to string
			book.ID = params["id"] // Just an example - Not Production Safe

			// Append newly edited book to collection
			books = append(books, book)

			// Add new book item to response
			json.NewEncoder(w).Encode(book)
			return
		}
		break
	}

	// Respond with new set of books
	json.NewEncoder(w).Encode(books)
}

// deleteBooks delete single book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Search through books and when match is found
	//
	for index, item := range books {
		if item.ID == params["id"] {
			// Take slice up until (not including index)
			// and append remaining list
			// Effectively removes specific book
			books = append(books[:index], books[index+1:]...)
			break
		}
		break
	}

	// Respond with new set of books
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init Router (using type inference)
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	// Appending objects to books slice
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "12345", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route Handlers (aka Endpoints)
	// GET all books
	r.HandleFunc("/api/books", getBooks).Methods("GET")

	// GET single book
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")

	// POST create book
	r.HandleFunc("/api/books", createBook).Methods("POST")

	// PUT edit book
	r.HandleFunc("/api/books/{id}", updateBook).Methods("GET")

	// DELETE remove book
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Similar to app.serve() in React
	// Pass in desired port and router
	// log.Fatal() to log errors
	log.Fatal(http.ListenAndServe(":8000", r))
}
