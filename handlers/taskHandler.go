package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/hawkerd/jira-clone/models"
)

var tasks []models.Task // list of tasks
var taskID int = 0

// '/tasks'
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		listTasksHandler(w, r)
	} else if r.Method == http.MethodPost {
		createTaskHandler(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Only GET and POST methods are allowed")
	}
}

// '/tasks' GET
func listTasksHandler(w http.ResponseWriter, r *http.Request) {
	for _, task := range tasks {
		fmt.Fprintf(w, "Task: %d, Title: %s, Status: %s\n", task.ID, task.Title, task.Status)
	}
}

// '/tasks' POST
func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	status := r.FormValue("status")

	newTask := models.Task{
		ID:          taskID,
		Description: description,
		Status:      status,
		Title:       title,
	}

	taskID++
	tasks = append(tasks, newTask)

	fmt.Fprintf(w, "Created Task: %d, Title: %s\n", newTask.ID, newTask.Title)
}

// '/tasks/{id}'
func TaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	// extract task id
	var pathParts = strings.Split(r.URL.Path, "/") // list of strings
	if len(pathParts) < 3 {                        // error if incorrectly formatted
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid task ID")
		return
	}
	id, err := strconv.Atoi(pathParts[2]) // error if we cant convert the id to an int
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid task ID")
		return
	}

	switch r.Method { //handle DELETE, PUT, GET
	case http.MethodGet:
		//

	case http.MethodDelete: // delete the specified task
		// verify that the task exists and is not deleted
		if (id < 1) || (id > taskID) || (tasks[id-1].Status == "Deleted") {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid task ID: %d\n", id)
			return
		}

		// delete task
		tasks[id-1].Status = "Deleted"

	case http.MethodPut:

	}
}
