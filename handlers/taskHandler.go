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
	switch r.Method { // handle GET, POST
	case http.MethodGet:
		// list tasks
		for _, task := range tasks {
			fmt.Fprintf(w, "Task: %d, Title: %s, Status: %s, Project ID: %d\n", task.ID, task.Title, task.Status, task.ProjectID)
		}
	case http.MethodPost:
		//extract title, description, status fields
		title := r.FormValue("title")
		description := r.FormValue("description")
		status := r.FormValue("status")
		projectIdString := r.FormValue("projectID")

		// convert project id to an integer
		projectId, err := strconv.Atoi(projectIdString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Invalid project ID")
			return
		}

		projectExists := false
		for _, project := range projects {
			if project.ID == projectId {
				projectExists = true
				break
			}
		}

		if !projectExists {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Project not found")
			return
		}

		// create new task
		taskID++
		newTask := models.Task{
			ID:          taskID,
			Description: description,
			Status:      status,
			Title:       title,
			ProjectID:   projectId,
		}
		tasks = append(tasks, newTask)

		fmt.Fprintf(w, "Task created: %d, Title: %s, Status: %s, Project ID: %d\n", newTask.ID, newTask.Title, newTask.Status, newTask.ProjectID)
	}
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
		// verify that the task exists and is not deleted
		if (id < 1) || (id > taskID) || (tasks[id-1].Status == "Deleted") {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Invalid task ID: %d\n", id)
			return
		}

		// write task information
		task := tasks[id-1]
		fmt.Fprintf(w, "Task ID: %d\nTitle: %s\nDescription: %s\nStatus: %s",
			task.ID, task.Title, task.Description, task.Status)

	case http.MethodDelete: // delete the specified task
		// verify that the task exists and is not deleted
		if (id < 1) || (id > taskID) || (tasks[id-1].Status == "Deleted") {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Invalid task ID: %d\n", id)
			return
		}

		// delete task
		tasks[id-1].Status = "Deleted"

	case http.MethodPut:
		// verify that the task exists and is not deleted
		if (id < 1) || (id > taskID) || (tasks[id-1].Status == "Deleted") {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Invalid task ID: %d\n", id)
			return
		}

		task := &tasks[id-1]

		title := r.FormValue("title")
		description := r.FormValue("description")
		status := r.FormValue("status")

		// if the field was provided, update the task
		if title != "" {
			task.Title = title
		}
		if description != "" {
			task.Description = description
		}
		if status != "" {
			task.Status = status
		}

		fmt.Fprintf(w, "Updated Task: %d\nTitle: %s\nDescription: %s\nStatus: %s\n",
			task.ID, task.Title, task.Description, task.Status)

	}
}
