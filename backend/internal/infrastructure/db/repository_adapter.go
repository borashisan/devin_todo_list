package db

import (
	"backend/internal/domain"
	"context"
)

// TodoRepositoryAdapter adapts the sqlc-based TodoRepository to the domain.TodoRepository interface
type TodoRepositoryAdapter struct {
	repo *TodoRepository
}

// NewTodoRepositoryAdapter creates a new adapter
func NewTodoRepositoryAdapter(repo *TodoRepository) *TodoRepositoryAdapter {
	return &TodoRepositoryAdapter{repo: repo}
}

// List returns all todos
func (a *TodoRepositoryAdapter) List(ctx context.Context) ([]domain.Todo, error) {
	todos, err := a.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return toDomainTodos(todos), nil
}

// GetByID returns a todo by ID
func (a *TodoRepositoryAdapter) GetByID(ctx context.Context, id string) (*domain.Todo, error) {
	todo, err := a.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, nil
	}
	return toDomainTodo(todo), nil
}

// Create creates a new todo
func (a *TodoRepositoryAdapter) Create(ctx context.Context, title string) (*domain.Todo, error) {
	todo, err := a.repo.Create(ctx, title)
	if err != nil {
		return nil, err
	}
	return toDomainTodo(todo), nil
}

// Update updates a todo
func (a *TodoRepositoryAdapter) Update(ctx context.Context, id string, title string, isCompleted bool) (*domain.Todo, error) {
	todo, err := a.repo.Update(ctx, id, title, isCompleted)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, nil
	}
	return toDomainTodo(todo), nil
}

// Delete deletes a todo
func (a *TodoRepositoryAdapter) Delete(ctx context.Context, id string) error {
	return a.repo.Delete(ctx, id)
}

// toDomainTodo converts a db.Todo to domain.Todo
func toDomainTodo(t *Todo) *domain.Todo {
	return &domain.Todo{
		ID:          t.ID,
		Title:       t.Title,
		IsCompleted: t.IsCompleted,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

// toDomainTodos converts a slice of db.Todo to domain.Todo
func toDomainTodos(todos []Todo) []domain.Todo {
	result := make([]domain.Todo, len(todos))
	for i, t := range todos {
		result[i] = domain.Todo{
			ID:          t.ID,
			Title:       t.Title,
			IsCompleted: t.IsCompleted,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		}
	}
	return result
}
