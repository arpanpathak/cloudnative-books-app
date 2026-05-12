package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var (
	rdb *redis.ClusterClient
	ctx = context.Background()
)

func initRedis() {
	// Dynamically get the Redis connection details from environment variables
	// In Kubernetes, REDIS_ADDRS will be: redis-cluster.redis.svc.cluster.local:6379
	addrsStr := os.Getenv("REDIS_ADDRS")
	if addrsStr == "" {
		addrsStr = "localhost:6379" // Fallback for local testing
	}
	addrs := strings.Split(addrsStr, ",")

	password := os.Getenv("REDIS_PASSWORD")

	rdb = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addrs,
		Password: password,
	})

	// Test the connection
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Printf("⚠️ Redis Cluster ping failed: %v", err)
		log.Println("Starting anyway (will retry connections dynamically).")
	} else {
		log.Println("✅ Successfully connected to Redis Cluster")
	}
}

func main() {
	initRedis()

	http.HandleFunc("/books", booksHandler)
	http.HandleFunc("/books/", bookDetailHandler)

	log.Println("🚀 Starting books-api server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// Generate an ID. We use {books} hash tag to ensure the ID counter 
// and the collection hash sit on the exact same cluster shard.
func getNextID() (int, error) {
	id, err := rdb.Incr(ctx, "{books}:next_id").Result()
	return int(id), err
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// Store books in a Redis Hash at key "{books}:collection"
		booksMap, err := rdb.HGetAll(ctx, "{books}:collection").Result()
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		var list []Book
		for _, bookJSON := range booksMap {
			var b Book
			json.Unmarshal([]byte(bookJSON), &b)
			list = append(list, b)
		}
		json.NewEncoder(w).Encode(list)

	case http.MethodPost:
		var b Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		
		id, err := getNextID()
		if err != nil {
			http.Error(w, "Failed to generate ID", http.StatusInternalServerError)
			return
		}
		b.ID = id

		bookJSON, _ := json.Marshal(b)
		if err := rdb.HSet(ctx, "{books}:collection", strconv.Itoa(id), bookJSON).Err(); err != nil {
			http.Error(w, "Failed to save book", http.StatusInternalServerError)
			return
		}
		
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(b)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func bookDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	
	switch r.Method {
	case http.MethodGet:
		bookJSON, err := rdb.HGet(ctx, "{books}:collection", idStr).Result()
		if err == redis.Nil {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(bookJSON))

	case http.MethodPut:
		exists, err := rdb.HExists(ctx, "{books}:collection", idStr).Result()
		if err != nil || !exists {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		var updatedBook Book
		if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		id, _ := strconv.Atoi(idStr)
		updatedBook.ID = id
		bookJSON, _ := json.Marshal(updatedBook)

		if err := rdb.HSet(ctx, "{books}:collection", idStr, bookJSON).Err(); err != nil {
			http.Error(w, "Failed to update book", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(updatedBook)

	case http.MethodDelete:
		deleted, err := rdb.HDel(ctx, "{books}:collection", idStr).Result()
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
