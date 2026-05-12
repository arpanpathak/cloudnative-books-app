package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/arpanpathak/books-api/internal/models"
	"github.com/arpanpathak/books-api/internal/repository"
	"github.com/redis/go-redis/v9"
)

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		list, err := repository.GetAllBooks()
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		if list == nil {
			list = []models.Book{} // return empty array instead of null
		}
		json.NewEncoder(w).Encode(list)

	case http.MethodPost:
		var b models.Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		id, err := repository.GetNextID()
		if err != nil {
			http.Error(w, "Failed to generate ID", http.StatusInternalServerError)
			return
		}
		b.ID = id

		if err := repository.SaveBook(id, b); err != nil {
			http.Error(w, "Failed to save book", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(b)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func BookDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")

	switch r.Method {
	case http.MethodGet:
		book, err := repository.GetBook(idStr)
		if err == redis.Nil {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(book)

	case http.MethodPut:
		exists, err := repository.BookExists(idStr)
		if err != nil || !exists {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		var updatedBook models.Book
		if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		id, _ := strconv.Atoi(idStr)
		updatedBook.ID = id

		if err := repository.SaveBook(id, updatedBook); err != nil {
			http.Error(w, "Failed to update book", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(updatedBook)

	case http.MethodDelete:
		deleted, err := repository.DeleteBook(idStr)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		if deleted == 0 {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
