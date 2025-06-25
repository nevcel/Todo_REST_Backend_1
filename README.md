# todo-backend
This repository contains a REST API backend for managing todos. The backend is based on the [HttpRouter](https://github.com/julienschmidt/httprouter). The project is structured based on the well-known MVC pattern and uses the repository architectural design pattern for abstracting the data access logic.

# Description
## Data model
### Todo
| Field name  | Data type |
|-------------|-----------|
| Id          | string    |
| Title       | string    |
| Description | string    |
| Terminated  | bool      |

### JsonExtendedResponse
| Field name | Data type   |
|------------|-------------|
| Meta       | interface{} |
| Data       | interface{} |

### JsonDataResponse
| Field name | Data type |
|------------|-----------|
| Data       | []Todo    |

### ApiError
| Field name | Data type |
|------------|-----------|
| Status     | int       |
| Title      | string    |

### JsonErrorResponse
| Field name  | Data type |
|-------------|-----------|
| Error       | ApiError  |

## Features
The following endpoints are implemented:

| No. | HTTP Verb | Path              | Expects (JSON) | Returns (JSON)                 | HTTP Status                                           | Description         |
|-----|-----------|-------------------|----------------|--------------------------------|-------------------------------------------------------|---------------------|
| 1   | GET       | /api/v1           | Nothing        | Welcome string                 | 200 (success)                                         | Welcome string      |
| 2   | GET       | /api/v1/todos     | Nothing        | An array with todo entries     | 200 (success)                                         | Get a list of todos |
| 3   | GET       | /api/v1/todos/:id | Nothing        | The todo with the specified ID | 200 (success) or 404 (not found)                      | Get todo by ID      |
| 4   | POST      | /api/v1/todos     | A todo entry   | The new todo entry             | 201 (created) or 400 (Bad Request)                    | Create new todo     |
| 5   | PUT       | /api/v1/todos/:id | A todo entry   | The updated todo entry         | 200 (success) or 400 (Bad Request) or 404 (not found) | Update todo by ID   |

# Installation
Cross-plattform executable, which can be built to target platform using go sdk:
````
go build .
````

# Configuration
The project uses an `.env` file for the configuration of the adjustable variables.
Currently, the following variables can be set:

| No. | Variable name   | Allowed values |
|-----|-----------------|----------------|
| 1   | REPOSITORY_MODE | "mem", "csv"   |
| 2   | PORT            | 0-65535        |

# Usage
The project can be built to desired platform and then the backend runs on port 8080, listening on any interface available at the place of execution.

# License
MIT

# Authors
be