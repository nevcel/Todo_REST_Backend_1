// Package controllers contains the todo backend controllers
package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"path"
	"todo-rest-backend/models"
	"todo-rest-backend/models/configuration"
	"todo-rest-backend/models/repositories/factory"
	"todo-rest-backend/models/todo"
)

// UriBasePath uri base path
const UriBasePath = "/api"

// UriVersion uri version
const UriVersion = "v1"

// UriRessourceTodos uri ressource todos
const UriRessourceTodos = "/todos"

// UriRessourceTodosPathParameterName uri ressource todos path parameter name
const UriRessourceTodosPathParameterName = "{id}"

// GeneralErrorMessage general error message
const GeneralErrorMessage = "an error has occurred"

// Run does the running of the web server
func Run() error {
	repositoryInstance, err := factory.GetTodoRepositoryInstance()
	if err != nil {
		return err
	}
	err = models.SetTodoRepository(repositoryInstance)
	if err != nil {
		return err
	}
	err = models.Initialize()
	if err != nil {
		return err
	}

	backendHostUrl, err := configuration.GetBackendHostUrl()
	if err != nil {
		return err
	}

	// StrictSlash == true: if the route path is "/path/", then a redirect to the path "/path" is done.
	router := mux.NewRouter().StrictSlash(true)

	api := router.PathPrefix(path.Join(UriBasePath, UriVersion)).Subrouter()
	api.HandleFunc("", Index).Methods("GET")
	api.HandleFunc(UriRessourceTodos, TodosGet).Methods("GET")
	api.HandleFunc(path.Join(UriRessourceTodos, UriRessourceTodosPathParameterName), TodoGetById).Methods("GET")
	api.HandleFunc(UriRessourceTodos, TodoPost).Methods("POST")
	api.HandleFunc(path.Join(UriRessourceTodos, UriRessourceTodosPathParameterName), TodoPut).Methods("PUT")

	fmt.Println("Backend running at:", backendHostUrl)
	err = http.ListenAndServe(backendHostUrl, router)
	if err != nil {
		return err
	}

	return nil
}

// Index Handler for the index action
// GET /
func Index(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(writer, "Welcome to the Todo REST API %s!\n", UriVersion)
	if err != nil {
		panic(err)
	}
}

// TodosGet Handler for the todos get action
// GET /todos
func TodosGet(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	todos, err := models.ReadTodos()
	if err != nil {
		handleErrorAndDiscloseDetails(writer, http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	sortedTodos := models.SortTodosAfterIdAscending(todos)
	response := models.JsonDataResponse{Data: sortedTodos}
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		panic(err)
	}
}

func handleErrorAndDiscloseDetails(writer http.ResponseWriter, statusCode int) {
	errorText := getGeneralErrorTextToAvoidDisclosingDetails()
	handleError(writer, statusCode, errorText)
}

func getGeneralErrorTextToAvoidDisclosingDetails() string {
	return GeneralErrorMessage
}

func handleError(writer http.ResponseWriter, statusCode int, text string) {
	writer.WriteHeader(statusCode)
	response := models.JsonErrorResponse{
		Error: models.ApiError{
			Status: statusCode,
			Title:  text,
		},
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		panic(err)
	}
}

// TodoGetById Handler for a todo get by id action
func TodoGetById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// Get id from url parameters
	vars := mux.Vars(request)
	id := vars["id"]
	todoRead, err := models.ReadTodoById(id)
	if err != nil {
		handleErrorAndDiscloseDetails(writer, http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := models.JsonExtendedResponse{Data: todoRead}
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		panic(err)
	}
}

// TodoPost Handler for the todos post action
func TodoPost(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var todoToCreate todo.Todo
	err := decodeTodo(request, &todoToCreate)
	if err != nil {
		handleError(writer, http.StatusBadRequest, err.Error())
		return
	}

	todoAdded, err := models.CreateTodo(todoToCreate)
	if err != nil {
		handleErrorAndDiscloseDetails(writer, http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	response := models.JsonExtendedResponse{Data: todoAdded}
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		panic(err)
	}
}

// decodeTodo decodes the json request body into a Todo
func decodeTodo(request *http.Request, todo *todo.Todo) error {
	if request.Body == nil {
		return errors.New("invalid body")
	}
	err := json.NewDecoder(request.Body).Decode(todo)
	if err != nil {
		return err
	}
	if todo.Title == "" || todo.Description == "" {
		err = errors.New("body: required fields missing")
		return err
	}
	return nil
}

// TodoPut Handler for a todo put by id action
func TodoPut(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// Get id from url parameters
	vars := mux.Vars(request)
	id := vars["id"]

	// Get update todo from request body
	var todoToUpdate todo.Todo
	err := decodeTodo(request, &todoToUpdate)
	if err != nil {
		handleErrorAndDiscloseDetails(writer, http.StatusBadRequest)
		return
	}

	todoUpdated, err := models.UpdateTodoById(id, todoToUpdate)
	if err != nil {
		handleErrorAndDiscloseDetails(writer, http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	response := models.JsonExtendedResponse{Data: todoUpdated}
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		panic(err)
	}
}
