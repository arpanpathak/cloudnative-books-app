package repository

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/arpanpathak/books-api/internal/models"
	"github.com/redis/go-redis/v9"
)

var (
	RDB *redis.ClusterClient
	ctx = context.Background()
)

// InitRedis connects to the Redis cluster
func InitRedis() {
	addrsStr := os.Getenv("REDIS_ADDRS")
	if addrsStr == "" {
		addrsStr = "localhost:6379"
	}
	addrs := strings.Split(addrsStr, ",")
	password := os.Getenv("REDIS_PASSWORD")

	RDB = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addrs,
		Password: password,
	})

	if err := RDB.Ping(ctx).Err(); err != nil {
		log.Printf("⚠️ Redis Cluster ping failed: %v", err)
		log.Println("Starting anyway (will retry connections dynamically).")
	} else {
		log.Println("✅ Successfully connected to Redis Cluster")
	}
}

func GetNextID() (int, error) {
	id, err := RDB.Incr(ctx, "{books}:next_id").Result()
	return int(id), err
}

func GetAllBooks() ([]models.Book, error) {
	booksMap, err := RDB.HGetAll(ctx, "{books}:collection").Result()
	if err != nil {
		return nil, err
	}

	var list []models.Book
	for _, bookJSON := range booksMap {
		var b models.Book
		json.Unmarshal([]byte(bookJSON), &b)
		list = append(list, b)
	}
	return list, nil
}

func GetBook(id string) (models.Book, error) {
	var b models.Book
	bookJSON, err := RDB.HGet(ctx, "{books}:collection", id).Result()
	if err != nil {
		return b, err
	}
	json.Unmarshal([]byte(bookJSON), &b)
	return b, nil
}

func SaveBook(id int, b models.Book) error {
	bookJSON, _ := json.Marshal(b)
	return RDB.HSet(ctx, "{books}:collection", strconv.Itoa(id), bookJSON).Err()
}

func DeleteBook(id string) (int64, error) {
	return RDB.HDel(ctx, "{books}:collection", id).Result()
}

func BookExists(id string) (bool, error) {
	return RDB.HExists(ctx, "{books}:collection", id).Result()
}
