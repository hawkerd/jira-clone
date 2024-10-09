package main

import (
	"fmt"
	"net/http"

	"github.com/hawkerd/jira-clone/database"
	"github.com/hawkerd/jira-clone/handlers"
	"github.com/rs/cors"
)

func main() {
	// initialize connection to database
	database.Init()

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/tasks", handlers.TasksHandler)
	mux.HandleFunc("/tasks/", handlers.TaskByIDHandler)
	mux.HandleFunc("/projects", handlers.ProjectHandler)
	mux.HandleFunc("/projects/", handlers.ProjectByIDHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5500", "http://localhost:3000"}, // restrict to specific domain
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	fmt.Println("Starting server on port :8080")

	if err := http.ListenAndServe(":8080", handler); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
