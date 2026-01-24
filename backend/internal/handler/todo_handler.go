package handler

import (
	"backend/internal/domain"
	"backend/internal/usecase"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// TodoHandler handles HTTP requests for todos
type TodoHandler struct {
	usecase *usecase.TodoUsecase
}

// NewTodoHandler creates a new TodoHandler
func NewTodoHandler(usecase *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{usecase: usecase}
}

// CreateTodoRequest represents the request body for creating a todo
type CreateTodoRequest struct {
	Title string `json:"title"`
}

// UpdateTodoRequest represents the request body for updating a todo's completion status
type UpdateTodoRequest struct {
	IsCompleted bool `json:"is_completed"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// ListTodos handles GET /api/todos
func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	todos, err := h.usecase.List(ctx)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Return empty array instead of null when no todos
	if todos == nil {
		todos = []domain.Todo{}
	}

	h.respondJSON(w, http.StatusOK, todos)
}

// CreateTodo handles POST /api/todos
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Title == "" {
		h.respondError(w, http.StatusBadRequest, "Title is required")
		return
	}

	todo, err := h.usecase.Create(ctx, req.Title)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondJSON(w, http.StatusCreated, todo)
}

// UpdateTodoCompleted handles PATCH /api/todos/{id}
func (h *TodoHandler) UpdateTodoCompleted(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	var req UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	todo, err := h.usecase.UpdateCompleted(ctx, id, req.IsCompleted)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if todo == nil {
		h.respondError(w, http.StatusNotFound, "Todo not found")
		return
	}

	h.respondJSON(w, http.StatusOK, todo)
}

// DeleteTodo handles DELETE /api/todos/{id}
func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	if err := h.usecase.Delete(ctx, id); err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// respondJSON sends a JSON response
func (h *TodoHandler) respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// respondError sends an error response
func (h *TodoHandler) respondError(w http.ResponseWriter, status int, message string) {
	h.respondJSON(w, status, ErrorResponse{Error: message})
}
