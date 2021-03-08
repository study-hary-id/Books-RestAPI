package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
Book is used to store all of metadata from a book.
*/
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

/*
Author is used to store the name of the author of the book.
*/
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

/*
books is a collection of book.
*/
var books []Book

/*
getBook handler will give a specific book with particular id,
if there is no desired book then it will return blank Book type.
*/
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var param = mux.Vars(r)

	for _, book := range books {
		if book.ID == param["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

/*
getBooks handler will give all book data within the API.
*/
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

/*
createBook handler will add one book to the API or to the database.
*/
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		book        Book
		latestId, _ = strconv.Atoi(books[len(books)-1].ID)
	)

	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(latestId + 1)
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

/*
updateBook handler will update data from a book within the API or database.
If there is no desired book, it will response all data within the API.
*/
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		newBook     Book
		param       = mux.Vars(r)
		latestId, _ = strconv.Atoi(books[len(books)-1].ID)
	)

	for i, currBook := range books {
		if currBook.ID == param["id"] {

			books = append(books[:i], books[i+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&newBook)
			newBook.ID = strconv.Itoa(latestId + 1)
			books = append(books, newBook)

			json.NewEncoder(w).Encode(newBook)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

/*
deleteBook handler will delete a specific book with particular id.
*/
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var param = mux.Vars(r)

	for i, book := range books {
		if book.ID == param["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Initialization a new router
	var r = mux.NewRouter()

	// Hardcoded or dummy data for personal use
	books = append(books, Book{
		ID:    "1",
		ISBN:  "978-0-596-15990-0",
		Title: "Head First HTML and CSS Second Edition",
		Author: &Author{
			FirstName: "Elisabeth Robson",
			LastName:  "Eric Freeman",
		},
	})

	// Route handlers and some endpoints in this RESTful API
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// I don't know yet what it's used for, but more or less for starting the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
