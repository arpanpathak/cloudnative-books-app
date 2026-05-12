package main

import (
	"log"
	"net/http"

	"github.com/arpanpathak/books-api/internal/handlers"
	"github.com/arpanpathak/books-api/internal/repository"
)

func main() {
	// 1. Initialize dependencies
	repository.InitRedis()

	// 2. Register HTTP Routes
	http.HandleFunc("/books", handlers.BooksHandler)
	http.HandleFunc("/books/", handlers.BookDetailHandler)

	// 3. Start Server
	log.Println("🚀 Starting books-api server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
