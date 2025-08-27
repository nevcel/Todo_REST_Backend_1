// Package models contains the todo backend models and corresponding business logic
package models

import (
	"errors"
	"sort"
	"strconv"
	"todo-rest-backend/models/repositories"
	"todo-rest-backend/models/todo"
)

// JsonExtendedResponse type definition with json tags
type JsonExtendedResponse struct {
	Meta interface{} `json:"meta"` // Field to add some meta information to the API response
	Data interface{} `json:"data"`
}

// JsonDataResponse type definition with json tags
type JsonDataResponse struct {
	Data []todo.Todo `json:"data"`
}

// JsonErrorResponse type definition with json tags
type JsonErrorResponse struct {
	Error ApiError `json:"error"`
}

// ApiError type definition with json tags
type ApiError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

var todoRepository repositories.TodoRepository

// SetTodoRepository allows to set the repositories type
func SetTodoRepository(todoRepositoryNew repositories.TodoRepository) error {
	if todoRepositoryNew == nil {
		return errors.New("todo repositories must not be nil")
	}
	todoRepository = todoRepositoryNew
	return nil
}

// Initialize initializes the repository (abstracted by repository pattern)
func Initialize() error {
	if todoRepository == nil {
		return errors.New("todo repositories must not be nil")
	}
	return todoRepository.Initialize()
}

// ReadTodos returns todo's from repository (abstracted by repository pattern)
func ReadTodos() ([]todo.Todo, error) {
	if todoRepository == nil {
		return nil, errors.New("todo repositories must not be nil")
	}
	return todoRepository.ReadTodos()
}

// ReadTodoById returns todo with passed id when existing from repository (abstracted by repository pattern)
func ReadTodoById(id string) (todo.Todo, error) {
	if todoRepository == nil {
		return todo.Todo{}, errors.New("todo repositories must not be nil")
	}
	return todoRepository.ReadTodoById(id)
}

// CreateTodo stores the passed todo in the repository and returns the stored todo (abstracted by repository pattern)
func CreateTodo(todoToCreate todo.Todo) (todo.Todo, error) {
	if todoRepository == nil {
		return todo.Todo{}, errors.New("todoToCreate repositories must not be nil")
	}
	return todoRepository.CreateTodo(todoToCreate)
}

// SortTodosAfterIdAscending sorts the todos ascending after the id and returns sorted todos
func SortTodosAfterIdAscending(todos []todo.Todo) []todo.Todo {
	sort.Slice(todos, func(i, j int) bool {
		leftValueAsInt, _ := strconv.Atoi(todos[i].Id)
		rightValueAsInt, _ := strconv.Atoi(todos[j].Id)
		return leftValueAsInt < rightValueAsInt
	})
	return todos
}

// UpdateTodoById returns updated todo from repository (abstracted by repository pattern)
func UpdateTodoById(id string, todoUpdate todo.Todo) (todo.Todo, error) {
	if todoRepository == nil {
		return todo.Todo{}, errors.New("todo repositories must not be nil")
	}
	return todoRepository.UpdateTodoById(id, todoUpdate)
}

// DeleteTodoById delete todo from repository (abstracted by repository pattern)
func DeleteTodoById(id string, todoDelete todo.Todo) (todo.Todo, error) {
	if todoRepository == nil {
		return todo.Todo{}, errors.New("todoToDelete repositories must not be nil")
	}
	return todoRepository.DeleteTodoById(id, todoDelete)
}
