// Package memrepo contains the repository logic for memory I/O operations
package memrepo

import (
	"errors"
	"fmt"
	"strconv"
	"todo-rest-backend/models/todo"
)

// MemoryTodoRepository type
type MemoryTodoRepository struct {
	todoStore []todo.Todo
}

// Initialize initializes the repository
func (m *MemoryTodoRepository) Initialize() error {
	m.todoStore = []todo.Todo{}
	return nil
}

// ReadTodos returns todo's stored in memory
func (m *MemoryTodoRepository) ReadTodos() ([]todo.Todo, error) {
	return clone(m.todoStore), nil
}

func clone(todos []todo.Todo) []todo.Todo {
	var todoStoreClone []todo.Todo

	for _, currentTodo := range todos {
		todoStoreClone = append(todoStoreClone, currentTodo)
	}

	return todoStoreClone
}

// ReadTodoById returns todo stored in memory with passed id when existing
func (m *MemoryTodoRepository) ReadTodoById(id string) (todo.Todo, error) {
	for _, currentTodo := range m.todoStore {
		if id == currentTodo.Id {
			return currentTodo, nil
		}
	}

	return todo.Todo{}, errors.New("id not found")
}

// CreateTodo stores the passed todo in memory and returns the stored todo
func (m *MemoryTodoRepository) CreateTodo(todo todo.Todo) (todo.Todo, error) {
	todoCount := len(m.todoStore) + 1
	todo.Id = strconv.Itoa(todoCount)
	m.todoStore = append(m.todoStore, todo)

	return todo, nil
}

// UpdateTodoById updates the passed todo by id in memory and returns the updated todo
func (m *MemoryTodoRepository) UpdateTodoById(id string, todoUpdate todo.Todo) (todo.Todo, error) {
	for index, currentTodo := range m.todoStore {
		if currentTodo.Id == id {
			// update todo based on input
			todoUpdate.Id = id
			m.todoStore[index] = todoUpdate

			return todoUpdate, nil
		}
	}

	return todo.Todo{}, fmt.Errorf("item with id %s not found. Updating not possible", id)
}

func (m *MemoryTodoRepository) DeleteTodoById(id string, todoDelete todo.Todo) (todo.Todo, error) {
	return todoDelete, nil
}
