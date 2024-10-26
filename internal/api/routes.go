package application

import (
	"net/http"

	"github.com/bmena34/wordie/internal/auth"
	"github.com/bmena34/wordie/internal/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-redis/redis"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/word", func(r chi.Router) {
		r.Use(auth.AuthMiddleware)
		loadWordRoutes(r)
	})

	return router
}

func loadWordRoutes(router chi.Router) {
	rbd := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	wordHandler := &handler.Word{
		Rdb: rbd,
	}

	router.Get("/{id}", wordHandler.GetByID)
}
