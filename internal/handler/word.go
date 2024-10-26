package handler

import (
	"encoding/json"
	"log"
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
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	val, err := wd.Rdb.HGetAll(id).Result()
	if err == redis.Nil {
		http.Error(w, "word not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("failed to get word for id %s: %v", id, err)
		http.Error(w, "failed to get word", http.StatusInternalServerError)

		return
	}

	if len(val) == 0 {
		log.Printf("no data found for id %s", id)
		http.Error(w, "word not found", http.StatusNotFound)
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
