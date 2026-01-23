package db

import (
	"context"
	"database/sql"
	"fmt"
)

type TodoRepository struct {
	db      *sql.DB
	queries *Queries
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db:      db,
		queries: New(db),
	}
}

func (r *TodoRepository) GetDB() *sql.DB {
	return r.db
}

func (r *TodoRepository) GetQueries() *Queries {
	return r.queries
}

func (r *TodoRepository) Create(ctx context.Context, title string) (*Todo, error) {
	params := CreateTodoParams{
		Title:       title,
		IsCompleted: false,
	}

	result, err := r.queries.CreateTodo(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id: %w", err)
	}

	var id string
	err = r.db.QueryRowContext(ctx, "SELECT id FROM todos WHERE ROWID = ?", lastID).Scan(&id)
	if err != nil {
		todos, listErr := r.queries.ListTodos(ctx)
		if listErr != nil || len(todos) == 0 {
			return nil, fmt.Errorf("failed to retrieve created todo: %w", err)
		}
		return &todos[0], nil
	}

	return r.queries.GetTodo(ctx, id)
}

func (r *TodoRepository) GetByID(ctx context.Context, id string) (*Todo, error) {
	todo, err := r.queries.GetTodo(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get todo: %w", err)
	}
	return &todo, nil
}

func (r *TodoRepository) List(ctx context.Context) ([]Todo, error) {
	todos, err := r.queries.ListTodos(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list todos: %w", err)
	}
	return todos, nil
}

func (r *TodoRepository) Update(ctx context.Context, id string, title string, isCompleted bool) (*Todo, error) {
	params := UpdateTodoParams{
		ID:          id,
		Title:       title,
		IsCompleted: isCompleted,
	}

	result, err := r.queries.UpdateTodo(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update todo: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return nil, nil
	}

	return r.GetByID(ctx, id)
}

func (r *TodoRepository) Delete(ctx context.Context, id string) error {
	err := r.queries.DeleteTodo(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}
	return nil
}

func (r *TodoRepository) SearchByTitle(ctx context.Context, titlePattern string) ([]Todo, error) {
	todos, err := r.queries.GetTodoByTitle(ctx, titlePattern)
	if err != nil {
		return nil, fmt.Errorf("failed to search todos by title: %w", err)
	}
	return todos, nil
}

func ConnectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
