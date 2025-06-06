// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: todos.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (name, valid_till, completed, completed_at, created_at, updated_at)
VALUES (?1, ?2, ?3, ?4, ?5, ?6)
RETURNING id, name, created_at, updated_at, valid_till, completed, completed_at
`

type CreateTodoParams struct {
	Name        string
	ValidTill   time.Time
	Completed   bool
	CompletedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo,
		arg.Name,
		arg.ValidTill,
		arg.Completed,
		arg.CompletedAt,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ValidTill,
		&i.Completed,
		&i.CompletedAt,
	)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?1
`

func (q *Queries) DeleteTodo(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const getAllTodos = `-- name: GetAllTodos :many
SELECT id, name, created_at, updated_at, valid_till, completed, completed_at FROM todos
`

func (q *Queries) GetAllTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, getAllTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ValidTill,
			&i.Completed,
			&i.CompletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTodo = `-- name: GetTodo :one
SELECT id, name, created_at, updated_at, valid_till, completed, completed_at FROM todos
WHERE id = ?1
LIMIT 1
`

func (q *Queries) GetTodo(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ValidTill,
		&i.Completed,
		&i.CompletedAt,
	)
	return i, err
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos
SET name = ?1, valid_till = ?2, completed = ?3, completed_at = ?4, updated_at = ?5
WHERE id = ?6
RETURNING id, name, created_at, updated_at, valid_till, completed, completed_at
`

type UpdateTodoParams struct {
	Name        string
	ValidTill   time.Time
	Completed   bool
	CompletedAt sql.NullTime
	UpdatedAt   time.Time
	ID          int64
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodo,
		arg.Name,
		arg.ValidTill,
		arg.Completed,
		arg.CompletedAt,
		arg.UpdatedAt,
		arg.ID,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ValidTill,
		&i.Completed,
		&i.CompletedAt,
	)
	return i, err
}
