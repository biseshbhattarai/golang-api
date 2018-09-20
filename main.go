package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Book struct (Model)
type Book struct {
	gorm.Model
	Name   string
	Pages  string
	Author string
}

// Init books var as a slice Book struct
var books []Book

// Get all books

// Get single book
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := gorm.Open("sqlite3", "book.db")
	if err != nil {
		panic("Couldn't connect to the database")
	}
	defer db.Close()
	db.Find(&books)
	json.NewEncoder(w).Encode(books)
}

// Add new book
func createBook(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "book.db")
	if err != nil {
		panic("Couldn't connect to the database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	name := vars["name"]
	page := vars["page"]

	author := vars["author"]
	db.Create(&Book{Name: name, Pages: page, Author: author})
	fmt.Fprintf(w, "New Book successfully added")
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "book.db")
	if err != nil {
		panic("Couldn't connect to the database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	name := vars["name"]
	author := vars["author"]
	var book Book
	db.Where("name = ?", name).Find(&book)
	book.Author = author
	db.Save(&book)
	fmt.Fprintf(w, "Succesfully updated")
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "book.db")
	if err != nil {
		panic("Couldn't connect to the database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	name := vars["name"]

	var book Book
	db.Where("name = ?", name).Find(&book)
	db.Delete(&book)
	fmt.Fprintf(w, "Succesfully deleted")
}

func requestRoutes() {
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))

}

func initialMigration() {
	db, err := gorm.Open("sqlite3", "book.db")
	if err != nil {
		panic("Couldn't connect to database")
	}
	defer db.Close()
	db.AutoMigrate(&Book{})

}

// Main function
func main() {
	initialMigration()
	requestRoutes()

}
