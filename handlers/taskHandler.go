package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/hawkerd/jira-clone/database"
	"github.com/hawkerd/jira-clone/models"
)

// '/tasks'
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method { // handle GET, POST
	case http.MethodGet:
		// retrieve parameters
		projectIdStr := r.URL.Query().Get("projectID")
		statusFilter := r.URL.Query().Get("status")

		// query the database
		var filteredTasks []models.Task
		query := database.DB

		// apply project filter
		if projectIdStr != "" {
			projectId, err := strconv.Atoi(projectIdStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Invalid project ID")
				return
			} else {
				query = query.Where("project_id = ?", projectId)
			}
		}

		// apply status filter
		if statusFilter != "" {
			query = query.Where("status = ?", statusFilter)
		}

		// execute query
		query.Find(&filteredTasks)

		// list tasks
		if len(filteredTasks) == 0 {
			fmt.Fprintln(w, "No tasks found")
			return
		} else {
			for _, task := range filteredTasks {
				fmt.Fprintf(w, "Task: %d, Title: %s, Status: %s, Project ID: %d\n", task.ID, task.Title, task.Status, task.ProjectID)
			}
		}

	case http.MethodPost:
		// extract title, description, status fields
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

		// test whether the project exists
		var project models.Project
		err = database.DB.First(&project, projectId).Error
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Project not found")
			return
		}

		// add new task to the database
		newTask := models.Task{
			Description: description,
			Status:      status,
			Title:       title,
			ProjectID:   uint(projectId),
		}
		database.DB.Create(&newTask)

		// return info about the task created
		fmt.Fprintf(w, "Task created: %d, Title: %s, Status: %s, Project ID: %d\n", newTask.ID, newTask.Title, newTask.Status, newTask.ProjectID)
	}
}

// '/tasks/{id}'
func TaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	// extract task id
	var pathParts = strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid task ID")
		return
	}
	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid task ID")
		return
	}

	// verify that the task exists
	var task models.Task
	if err = database.DB.First(&task, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Task not found")
		return
	}

	switch r.Method { //handle DELETE, PUT, GET
	case http.MethodGet:
		// respond with task info
		fmt.Fprintf(w, "Task ID: %d\nTitle: %s\nDescription: %s\nStatus: %s\n",
			task.ID, task.Title, task.Description, task.Status)

	case http.MethodDelete: // delete the specified task
		// delete task (soft database delete)
		database.DB.Delete(&task)

		// respond with details about deleted task
		fmt.Fprintf(w, "Deleted task: %d\n", task.ID)

	case http.MethodPut:
		// update fields
		if title := r.FormValue("title"); title != "" {
			task.Title = title
		}
		if description := r.FormValue("description"); description != "" {
			task.Description = description
		}
		if status := r.FormValue("status"); status != "" {
			task.Status = status
		}

		// respond with info about the updated task
		fmt.Fprintf(w, "Updated Task: %d\nTitle: %s\nDescription: %s\nStatus: %s\n",
			task.ID, task.Title, task.Description, task.Status)

	}
}
