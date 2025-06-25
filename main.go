package main

import (
	"log"
	"todo-rest-backend/controllers"
)

func main() {
	err := controllers.Run()
	if err != nil {
		log.Fatal(err)
	}
}
