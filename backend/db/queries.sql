-- name: GetTodo :one
SELECT id, title, is_completed, created_at, updated_at
FROM todos
WHERE id = ?;

-- name: ListTodos :many
SELECT id, title, is_completed, created_at, updated_at
FROM todos
ORDER BY created_at DESC;

-- name: CreateTodo :execresult
INSERT INTO todos (id, title, is_completed)
VALUES (UUID(), ?, ?);

-- name: UpdateTodo :execresult
UPDATE todos
SET title = ?, is_completed = ?
WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;

-- name: GetTodoByTitle :many
SELECT id, title, is_completed, created_at, updated_at
FROM todos
WHERE title LIKE ?
ORDER BY created_at DESC;
