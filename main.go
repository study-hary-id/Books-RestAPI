package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

func getBook(w http.ResponseWriter, r *http.Request) {}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {}
func updateBook(w http.ResponseWriter, r *http.Request) {}
func deleteBook(w http.ResponseWriter, r *http.Request) {}

func main() {
	var r = mux.NewRouter()

	books = append(books, Book{
		ID:    "1",
		ISBN:  "978-0-596-15990-0",
		Title: "Head First HTML and CSS Second Edition",
		Author: &Author{
			FirstName: "Elisabeth Robson",
			LastName:  "Eric Freeman",
		},
	})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
