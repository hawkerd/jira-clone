package main

import (
	"fmt"
	"net/http"

	"github.com/hawkerd/jira-clone/handlers"
)

// '/'
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to your Jira clone!")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/tasks", handlers.TasksHandler)
	http.HandleFunc("/tasks/", handlers.TaskByIDHandler)
	http.HandleFunc("/projects", handlers.ProjectHandler)

	fmt.Println("Starting server on port :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
