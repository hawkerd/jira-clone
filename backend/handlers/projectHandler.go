package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hawkerd/jira-clone/database"
	"github.com/hawkerd/jira-clone/models"
)

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		// ensure name was provided
		name := r.FormValue("name")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Project name is required"})
			return
		}

		// create new project
		newProject := models.Project{
			Name: name,
		}

		// save the project to the database
		if err := database.DB.Create(&newProject).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create project"})
			return
		}

		// send project json
		json.NewEncoder(w).Encode(newProject)

	case http.MethodGet:
		var projects []models.Project

		// retrieve all projects from the database
		if err := database.DB.Find(&projects).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to load projects"})
			return
		}

		// send projects json
		json.NewEncoder(w).Encode(projects)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
	}
}

func ProjectByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// extract project id
	var pathParts = strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid project ID"})
		return
	}
	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid project ID"})
		return
	}

	// verify that the task exists
	var project models.Project
	if err = database.DB.First(&project, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Project not found"})
		return
	}

	switch r.Method {
	case http.MethodDelete: // delete the specified project
		// delete all tasks associated with the project
		if err := database.DB.Where("project_id = ?", id).Delete(&models.Task{}).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete tasks associated tasks"})
			return
		}

		// delete project (soft database delete)
		if err := database.DB.Delete(&project).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete project"})
			return
		}

		// respond with task json
		json.NewEncoder(w).Encode(map[string]string{"message": "Project and associated tasks deleted successfully"})
	}
}
