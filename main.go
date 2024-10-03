package main

import (
	"fmt"
	"net/http"

	"github.com/hawkerd/jira-clone/models"
)

var tasks []models.Task // list of tasks
var taskID int = 1

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to your Jira clone!")
}

func listTasksHandler(w http.ResponseWriter, r *http.Request) {
	for _, task := range tasks {
		fmt.Fprintf(w, "Task: %d, Title: %s, Status: %s\n", task.ID, task.Title, task.Status)
	}
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Only POST method is allowed")
		return
	}

	title := r.FormValue("title")
	description := r.FormValue("description")
	status := r.FormValue("status")

	newTask := models.Task{
		ID:          taskID,
		Description: description,
		Status:      status,
		Title:       title,
	}

	tasks = append(tasks, newTask)
	taskID++

	fmt.Fprintf(w, "Created Task: %d, Title: %s\n", newTask.ID, newTask.Title)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/tasks", listTasksHandler)
	http.HandleFunc("/tasks/create", createTaskHandler)

	fmt.Println("Starting server on port :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
