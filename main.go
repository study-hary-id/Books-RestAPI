package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

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
