// Package factory contains logic for creating todo repository instances
package factory

import (
	"todo-rest-backend/models/configuration"
	"todo-rest-backend/models/repositories"
	"todo-rest-backend/models/repositories/csvrepo"
	"todo-rest-backend/models/repositories/memrepo"
)

// GetTodoRepositoryInstance returns the configured todo repository instance (factory design pattern function)
func GetTodoRepositoryInstance() (repositories.TodoRepository, error) {
	repositoryMode, err := configuration.GetRepositoryMode()
	if err != nil {
		return nil, err
	}

	switch repositoryMode {
	case configuration.MemoryRepository:
		{
			return &memrepo.MemoryTodoRepository{}, nil
		}
	case configuration.CsvFileRepository:
		{
			return &csvrepo.CsvFileTodoRepository{}, nil
		}
	default:
		return &memrepo.MemoryTodoRepository{}, nil
	}
}
