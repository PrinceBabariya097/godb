-- name: CreateTodo :one
INSERT INTO todos (name, valid_till, completed, completed_at, created_at, updated_at)
VALUES (:name, :valid_till, :completed, :completed_at, :created_at, :updated_at)
RETURNING *;

-- name: GetAllTodos :many
SELECT * FROM todos;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = :id
LIMIT 1;

-- name: UpdateTodo :one
UPDATE todos
SET name = :name, valid_till = :valid_till, completed = :completed, completed_at = :completed_at, updated_at = :updated_at
WHERE id = :id
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = :id;