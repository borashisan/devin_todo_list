package domain

import (
	"context"
	"time"
)

// Todo represents a todo item entity
type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TodoRepository defines the interface for todo data access
type TodoRepository interface {
	List(ctx context.Context) ([]Todo, error)
	GetByID(ctx context.Context, id string) (*Todo, error)
	Create(ctx context.Context, title string) (*Todo, error)
	Update(ctx context.Context, id string, title string, isCompleted bool) (*Todo, error)
	Delete(ctx context.Context, id string) error
}
