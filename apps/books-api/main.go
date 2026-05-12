package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var (
	books  = make(map[int]Book)
	nextID = 1
	mu     sync.Mutex
)

func main() {
	// Pre-populate some data
	books[nextID] = Book{ID: nextID, Title: "Kubernetes in Action", Author: "Marko Luksa"}
	nextID++

	// Standard REST Collection-Oriented routing
	http.HandleFunc("/books", booksHandler)
	http.HandleFunc("/books/", bookDetailHandler)

	log.Println("Starting books-api server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// Handles /books (GET for list, POST for create)
func booksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		list := make([]Book, 0, len(books))
		for _, b := range books {
			list = append(list, b)
		}
		mu.Unlock()
		json.NewEncoder(w).Encode(list)

	case http.MethodPost:
		var b Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		mu.Lock()
		b.ID = nextID
		books[nextID] = b
		nextID++
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(b)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handles /books/{id} (GET, PUT, DELETE)
func bookDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Path[len("/books/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	book, exists := books[id]

	switch r.Method {
	case http.MethodGet:
		if !exists {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(book)

	case http.MethodPut:
		if !exists {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}
		var updatedBook Book
		if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		updatedBook.ID = id
		books[id] = updatedBook
		json.NewEncoder(w).Encode(updatedBook)

	case http.MethodDelete:
		if !exists {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}
		delete(books, id)
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
