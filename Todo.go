package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
	"text/tabwriter"
	"time"
)

type Todo struct {
	id           int
	name         string
	created_at   time.Time
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

func (t *Todos) AddTodo(name string, valid_till time.Time) {
	newTodo := Todo{
		id:           len(*t),
		name:         name,
		created_at:   time.Now(),
		updated_at:   time.Now(),
		valid_till:   valid_till,
		completed:    false,
		completed_at: nil,
	}

	*t = append(*t, newTodo)
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

func (t *Todos) PrintTodos(todos *Todos) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)

	headers := []string{"ID", "Name", "Created At", "Updated At", "Valid Till", "Completed", "Completed At"}

	formatRow := func(cells []string) string {
		return "| " + strings.Join(cells, " \t| ") + " \t|"
	}

	fmt.Fprintln(w, formatRow(headers))

	for _, todo := range *todos {
		cells := []string{
			fmt.Sprintf("%d", todo.id),
			todo.name,
			todo.created_at.Format("2006-01-02 15:04:05"),
			todo.updated_at.Format("2006-01-02 15:04:05"),
			todo.valid_till.Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%v", todo.completed),
			fmt.Sprintf("%v", todo.completed_at),
		}

		fmt.Fprintln(w, formatRow(cells))
	}

	w.Flush()

}
