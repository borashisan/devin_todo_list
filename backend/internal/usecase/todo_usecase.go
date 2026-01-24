package usecase

import (
	"backend/internal/domain"
	"context"
)

// TodoUsecase handles business logic for todos
type TodoUsecase struct {
	repo domain.TodoRepository
}

// NewTodoUsecase creates a new TodoUsecase
func NewTodoUsecase(repo domain.TodoRepository) *TodoUsecase {
	return &TodoUsecase{repo: repo}
}

// List returns all todos
func (u *TodoUsecase) List(ctx context.Context) ([]domain.Todo, error) {
	return u.repo.List(ctx)
}

// Create creates a new todo with the given title
func (u *TodoUsecase) Create(ctx context.Context, title string) (*domain.Todo, error) {
	return u.repo.Create(ctx, title)
}

// UpdateCompleted updates the completion status of a todo
func (u *TodoUsecase) UpdateCompleted(ctx context.Context, id string, isCompleted bool) (*domain.Todo, error) {
	// First get the existing todo to preserve the title
	existing, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, nil // Not found
	}

	return u.repo.Update(ctx, id, existing.Title, isCompleted)
}

// Delete deletes a todo by ID
func (u *TodoUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
