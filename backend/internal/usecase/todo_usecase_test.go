package usecase

import (
	"backend/internal/domain"
	"context"
	"errors"
	"testing"
	"time"
)

// mockTodoRepository is a mock implementation of domain.TodoRepository
type mockTodoRepository struct {
	todos       map[string]*domain.Todo
	listErr     error
	getByIDErr  error
	createErr   error
	updateErr   error
	deleteErr   error
	createCount int
}

func newMockRepo() *mockTodoRepository {
	return &mockTodoRepository{
		todos: make(map[string]*domain.Todo),
	}
}

func (m *mockTodoRepository) List(ctx context.Context) ([]domain.Todo, error) {
	if m.listErr != nil {
		return nil, m.listErr
	}
	result := make([]domain.Todo, 0, len(m.todos))
	for _, t := range m.todos {
		result = append(result, *t)
	}
	return result, nil
}

func (m *mockTodoRepository) GetByID(ctx context.Context, id string) (*domain.Todo, error) {
	if m.getByIDErr != nil {
		return nil, m.getByIDErr
	}
	todo, ok := m.todos[id]
	if !ok {
		return nil, nil
	}
	return todo, nil
}

func (m *mockTodoRepository) Create(ctx context.Context, title string) (*domain.Todo, error) {
	if m.createErr != nil {
		return nil, m.createErr
	}
	m.createCount++
	todo := &domain.Todo{
		ID:          "test-id",
		Title:       title,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	m.todos[todo.ID] = todo
	return todo, nil
}

func (m *mockTodoRepository) Update(ctx context.Context, id string, title string, isCompleted bool) (*domain.Todo, error) {
	if m.updateErr != nil {
		return nil, m.updateErr
	}
	todo, ok := m.todos[id]
	if !ok {
		return nil, nil
	}
	todo.Title = title
	todo.IsCompleted = isCompleted
	todo.UpdatedAt = time.Now()
	return todo, nil
}

func (m *mockTodoRepository) Delete(ctx context.Context, id string) error {
	if m.deleteErr != nil {
		return m.deleteErr
	}
	delete(m.todos, id)
	return nil
}

func TestTodoUsecase_UpdateCompleted(t *testing.T) {
	tests := []struct {
		name        string
		setupRepo   func(*mockTodoRepository)
		id          string
		isCompleted bool
		wantErr     bool
		wantNil     bool
		wantStatus  bool
	}{
		{
			name: "successfully toggle to completed",
			setupRepo: func(m *mockTodoRepository) {
				m.todos["1"] = &domain.Todo{
					ID:          "1",
					Title:       "Test Todo",
					IsCompleted: false,
				}
			},
			id:          "1",
			isCompleted: true,
			wantErr:     false,
			wantNil:     false,
			wantStatus:  true,
		},
		{
			name: "successfully toggle to incomplete",
			setupRepo: func(m *mockTodoRepository) {
				m.todos["1"] = &domain.Todo{
					ID:          "1",
					Title:       "Test Todo",
					IsCompleted: true,
				}
			},
			id:          "1",
			isCompleted: false,
			wantErr:     false,
			wantNil:     false,
			wantStatus:  false,
		},
		{
			name:        "returns nil when todo not found",
			setupRepo:   func(m *mockTodoRepository) {},
			id:          "nonexistent",
			isCompleted: true,
			wantErr:     false,
			wantNil:     true,
		},
		{
			name: "returns error when GetByID fails",
			setupRepo: func(m *mockTodoRepository) {
				m.getByIDErr = errors.New("database error")
			},
			id:          "1",
			isCompleted: true,
			wantErr:     true,
		},
		{
			name: "returns error when Update fails",
			setupRepo: func(m *mockTodoRepository) {
				m.todos["1"] = &domain.Todo{
					ID:          "1",
					Title:       "Test Todo",
					IsCompleted: false,
				}
				m.updateErr = errors.New("update error")
			},
			id:          "1",
			isCompleted: true,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := newMockRepo()
			tt.setupRepo(repo)
			usecase := NewTodoUsecase(repo)

			result, err := usecase.UpdateCompleted(context.Background(), tt.id, tt.isCompleted)

			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if tt.wantNil {
				if result != nil {
					t.Errorf("expected nil, got %+v", result)
				}
				return
			}

			if result == nil {
				t.Error("expected non-nil result, got nil")
				return
			}

			if result.IsCompleted != tt.wantStatus {
				t.Errorf("expected IsCompleted=%v, got %v", tt.wantStatus, result.IsCompleted)
			}
		})
	}
}

func TestTodoUsecase_Create(t *testing.T) {
	tests := []struct {
		name      string
		setupRepo func(*mockTodoRepository)
		title     string
		wantErr   bool
		wantTitle string
	}{
		{
			name:      "successfully creates todo",
			setupRepo: func(m *mockTodoRepository) {},
			title:     "New Todo",
			wantErr:   false,
			wantTitle: "New Todo",
		},
		{
			name: "returns error when Create fails",
			setupRepo: func(m *mockTodoRepository) {
				m.createErr = errors.New("create error")
			},
			title:   "New Todo",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := newMockRepo()
			tt.setupRepo(repo)
			usecase := NewTodoUsecase(repo)

			result, err := usecase.Create(context.Background(), tt.title)

			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result.Title != tt.wantTitle {
				t.Errorf("expected title=%q, got %q", tt.wantTitle, result.Title)
			}
		})
	}
}

func TestTodoUsecase_List(t *testing.T) {
	tests := []struct {
		name      string
		setupRepo func(*mockTodoRepository)
		wantErr   bool
		wantCount int
	}{
		{
			name:      "returns empty list",
			setupRepo: func(m *mockTodoRepository) {},
			wantErr:   false,
			wantCount: 0,
		},
		{
			name: "returns all todos",
			setupRepo: func(m *mockTodoRepository) {
				m.todos["1"] = &domain.Todo{ID: "1", Title: "Todo 1"}
				m.todos["2"] = &domain.Todo{ID: "2", Title: "Todo 2"}
			},
			wantErr:   false,
			wantCount: 2,
		},
		{
			name: "returns error when List fails",
			setupRepo: func(m *mockTodoRepository) {
				m.listErr = errors.New("list error")
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := newMockRepo()
			tt.setupRepo(repo)
			usecase := NewTodoUsecase(repo)

			result, err := usecase.List(context.Background())

			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if len(result) != tt.wantCount {
				t.Errorf("expected %d todos, got %d", tt.wantCount, len(result))
			}
		})
	}
}

func TestTodoUsecase_Delete(t *testing.T) {
	tests := []struct {
		name      string
		setupRepo func(*mockTodoRepository)
		id        string
		wantErr   bool
	}{
		{
			name: "successfully deletes todo",
			setupRepo: func(m *mockTodoRepository) {
				m.todos["1"] = &domain.Todo{ID: "1", Title: "Todo 1"}
			},
			id:      "1",
			wantErr: false,
		},
		{
			name: "returns error when Delete fails",
			setupRepo: func(m *mockTodoRepository) {
				m.deleteErr = errors.New("delete error")
			},
			id:      "1",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := newMockRepo()
			tt.setupRepo(repo)
			usecase := NewTodoUsecase(repo)

			err := usecase.Delete(context.Background(), tt.id)

			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
