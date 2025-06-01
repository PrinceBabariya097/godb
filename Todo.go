package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
	"text/tabwriter"
	"time"
	db "tm/internal/database"
)

type Todo struct {
	name         string
	updated_at   time.Time
	valid_till   time.Time
	completed    bool
	completed_at *time.Time
}

type Todos []Todo

func (t *Todos) ValidateIndex(index int) error {
	if index < 0 && index > len(*t) {
		return errors.New("invalid index")
	}

	return nil
}

func (apiConfig *apiConfig) AddTodo(name string, valid_till time.Time) {
	newTodo := db.CreateTodoParams{
		Name:        name,
		ValidTill:   valid_till,
		Completed:   false,
		CompletedAt: sql.NullTime{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	ctx := context.Background()

	apiConfig.DB.CreateTodo(ctx, newTodo)
}

func (t *Todos) UpdateTodo(index int, name string, valid_till time.Time) error {
	todos := *t
	error := todos.ValidateIndex(index)
	if error != nil {
		return error
	}

	(*t)[index].name = name
	(*t)[index].valid_till = valid_till
	(*t)[index].updated_at = time.Now()

	return nil
}

func (t *Todos) CompleteTodo(index int) error {
	todos := *t
	error := todos.ValidateIndex(index)
	now := time.Now()

	if error != nil {
		return error
	}

	todos[index].completed = true
	todos[index].completed_at = &now

	return nil
}

func (t *Todos) DeleteTodo(index int) error {
	todos := *t
	error := todos.ValidateIndex(index)

	if error != nil {
		return error
	}

	deletedTodo := slices.Delete(todos, index, index+1)
	*t = deletedTodo

	return nil
}

func (apiConfig *apiConfig) PrintTodos() error {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)

	headers := []string{"ID", "Name", "Created At", "Updated At", "Valid Till", "Completed", "Completed At"}

	formatRow := func(cells []string) string {
		return "| " + strings.Join(cells, " \t| ") + " \t|"
	}

	todos, err := apiConfig.DB.GetAllTodos(context.Background())

	if err != nil {
		return err
	}

	fmt.Fprintln(w, formatRow(headers))

	for _, todo := range todos {
		cells := []string{
			fmt.Sprintf("%d", todo.ID),
			todo.Name,
			todo.CreatedAt.Format("2006-01-02 15:04:05"),
			todo.UpdatedAt.Format("2006-01-02 15:04:05"),
			todo.ValidTill.Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%v", todo.Completed),
			fmt.Sprintf("%v", todo.CompletedAt),
		}

		fmt.Fprintln(w, formatRow(cells))
	}

	w.Flush()

	return nil
}
