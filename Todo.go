package main

import (
	"errors"
	"slices"
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
