package main

import (
	"fmt"
	"net/http"

	"github.com/hawkerd/jira-clone/database"
	"github.com/hawkerd/jira-clone/handlers"
)

func main() {

	database.Init()

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/tasks", handlers.TasksHandler)
	http.HandleFunc("/tasks/", handlers.TaskByIDHandler)
	http.HandleFunc("/projects", handlers.ProjectHandler)

	fmt.Println("Starting server on port :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
