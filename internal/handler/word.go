package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

type Word struct {
	Rdb      *redis.Client `json:"-"`
	ID       int           `json:"id"`
	Category string        `json:"category"`
	Word     string        `json:"word"`
}

func (wd *Word) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println("id:", id)
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	val, err := wd.Rdb.HGetAll(id).Result()
	if err == redis.Nil {
		http.Error(w, "word not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "failed to get word", http.StatusInternalServerError)
		return
	}

	word := Word{
		Category: val["category"],
		Word:     val["word"],
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(word)
}
