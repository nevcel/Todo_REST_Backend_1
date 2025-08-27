// Package csvrepo contains the repository logic for csv file I/O operations
package csvrepo

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"todo-rest-backend/models/todo"
	"todo-rest-backend/models/utils"
)

// FileName for storage
const FileName = "data.csv"

// CsvFileTodoRepository type
type CsvFileTodoRepository struct {
}

// Initialize initializes the repository
func (c CsvFileTodoRepository) Initialize() error {
	var err error = nil

	// Check if file exists, if not create it
	//
	_, err = os.Stat(FileName)

	if os.IsNotExist(err) {
		var file *os.File
		file, err = os.Create(FileName)
		if err != nil {
			return err
		}
		defer utils.CloseFileAndHandleError(file, &err)
	}
	return err
}

// ReadTodos returns todo's stored in file
func (c CsvFileTodoRepository) ReadTodos() ([]todo.Todo, error) {
	return readDataFromFile()
}

func readDataFromFile() ([]todo.Todo, error) {
	file, err := os.Open(FileName)
	if err != nil {
		return nil, err
	}

	defer utils.CloseFileAndHandleError(file, &err)

	var readTodos []todo.Todo
	csvReader := csv.NewReader(file)
	for {
		var records []string
		records, err = csvReader.Read()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return nil, err
		}
		readTodos = append(readTodos, parseTodoData(records))
	}

	return readTodos, err
}

// ReadTodoById returns todo with passed id when existing
func (c CsvFileTodoRepository) ReadTodoById(id string) (todo.Todo, error) {
	todos, err := c.ReadTodos()
	if err != nil {
		return todo.Todo{}, err
	}

	for _, currentTodo := range todos {
		if id == currentTodo.Id {
			return currentTodo, err
		}
	}

	err = errors.New("id not found")
	return todo.Todo{}, err
}

// CreateTodo stores the passed todo in the file and returns the stored todo
func (c CsvFileTodoRepository) CreateTodo(todoToCreate todo.Todo) (todo.Todo, error) {
	todoCount, err := utils.GetLineCount(FileName)
	if err != nil {
		return todo.Todo{}, err
	}

	file, err := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return todo.Todo{}, err
	}
	defer utils.CloseFileAndHandleError(file, &err)

	todoToCreate.Id = strconv.Itoa(todoCount + 1)

	writer := csv.NewWriter(file)
	err = writer.Write(todoToCreate.Serialize())
	if err != nil {
		return todo.Todo{}, err
	}

	writer.Flush()

	return todoToCreate, err
}

func parseTodoData(rec []string) todo.Todo {
	id := rec[0]
	title := rec[1]
	description := rec[2]
	terminated := utils.ToBool(rec[3])

	todoParsed := todo.Todo{Id: id, Title: title, Description: description, Terminated: terminated}
	return todoParsed
}

// UpdateTodoById updates the passed todo by id in csv and returns the updated todo
func (c *CsvFileTodoRepository) UpdateTodoById(id string, todoUpdate todo.Todo) (todo.Todo, error) {
	// Create todo slice based on file
	todos, err := c.ReadTodos()
	if err != nil {
		return todo.Todo{}, err
	}

	// Update todo in slice
	itemFound := false
	for index, currentTodo := range todos {
		if id == currentTodo.Id {
			todoUpdate.Id = id
			todos[index] = todoUpdate
			itemFound = true
			break
		}
	}

	if itemFound == false {
		return todo.Todo{}, fmt.Errorf("item with id %s not found. Updating not possible", id)
	}

	// Clear file content
	err = utils.ClearFile(FileName)
	if err != nil {
		return todo.Todo{}, err
	}

	// Create todos from updated slice in file
	for _, currentTodo := range todos {
		_, err = c.CreateTodo(currentTodo)
		if err != nil {
			return todo.Todo{}, err
		}
	}

	return todoUpdate, nil
}
func (c *CsvFileTodoRepository) DeleteTodoById(id string, _ todo.Todo) (todo.Todo, error) {
	// Todos aus Datei lesen
	todos, err := c.ReadTodos()
	if err != nil {
		return todo.Todo{}, fmt.Errorf("error reading todos: %w", err)
	}

	// Todo finden und speichern bevor es gelöscht wird
	var deletedTodo todo.Todo
	var remainingTodos []todo.Todo
	itemFound := false

	for _, currentTodo := range todos {
		if id == currentTodo.Id {
			deletedTodo = currentTodo
			itemFound = true
			continue
		}
		remainingTodos = append(remainingTodos, currentTodo)
	}

	if !itemFound {
		return todo.Todo{}, fmt.Errorf("todo with ID %s not found", id)
	}

	// Datei leeren
	if err := utils.ClearFile(FileName); err != nil {
		return todo.Todo{}, fmt.Errorf("error clearing file: %w", err)
	}

	// Verbleibende Todos zurückschreiben
	for _, t := range remainingTodos {
		// Verwende direkt die CreateTodo Methode um die verbleibenden Todos zu speichern
		if _, err := c.CreateTodo(t); err != nil {
			return todo.Todo{}, fmt.Errorf("error rewriting todos: %w", err)
		}
	}

	return deletedTodo, nil
}
