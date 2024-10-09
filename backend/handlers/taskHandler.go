package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hawkerd/jira-clone/database"
	"github.com/hawkerd/jira-clone/models"
)

// '/tasks'
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
				json.NewEncoder(w).Encode(map[string]string{"error": "Invalid project ID"})
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
		if err := query.Find(&filteredTasks).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to retrieve tasks"})
			return
		}

		// send tasks json
		json.NewEncoder(w).Encode(filteredTasks)

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
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid project ID"})
			return
		}

		// test whether the project exists
		var project models.Project
		if err := database.DB.First(&project, projectId).Error; err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Project not found"})
			return
		}

		// add new task to the database
		newTask := models.Task{
			Description: description,
			Status:      status,
			Title:       title,
			ProjectID:   uint(projectId),
		}

		if err := database.DB.Create(&newTask).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create task"})
			return
		}

		// return task json
		json.NewEncoder(w).Encode(newTask)
	}
}

// '/tasks/{id}'
func TaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// extract task id
	var pathParts = strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid task ID"})
		return
	}
	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid task ID"})
		return
	}

	// verify that the task exists
	var task models.Task
	if err = database.DB.First(&task, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Task not found"})
		return
	}

	switch r.Method { //handle DELETE, PUT, GET
	case http.MethodGet:
		// respond with task json
		json.NewEncoder(w).Encode(task)

	case http.MethodDelete: // delete the specified task
		// delete task (soft database delete)
		if err := database.DB.Delete(&task).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete task"})
			return
		}

		// respond with task json
		json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})

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

		if err := database.DB.Save(&task).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update task"})
			return
		}

		// respond with task json
		json.NewEncoder(w).Encode(task)

	}
}
