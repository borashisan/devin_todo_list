package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// NewRouter creates a new chi router with CORS middleware
func NewRouter(todoHandler *TodoHandler) http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// CORS configuration for localhost:3000 (Nuxt frontend)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Root endpoint
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World from Go Backend!"))
	})

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Route("/todos", func(r chi.Router) {
			r.Get("/", todoHandler.ListTodos)
			r.Post("/", todoHandler.CreateTodo)
			r.Patch("/{id}", todoHandler.UpdateTodoCompleted)
			r.Delete("/{id}", todoHandler.DeleteTodo)
		})
	})

	return r
}
