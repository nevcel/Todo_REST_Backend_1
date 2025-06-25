// Package repositories contains the repository definitions
package repositories

import "todo-rest-backend/models/todo"

// TodoRepository interface todo repository type (used for repository architectural pattern interface definition)
type TodoRepository interface {
	Initialize() error
	ReadTodos() ([]todo.Todo, error)
	ReadTodoById(string) (todo.Todo, error)
	CreateTodo(todo.Todo) (todo.Todo, error)
	UpdateTodoById(string, todo.Todo) (todo.Todo, error)
}
