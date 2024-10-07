package handlers

import (
	"encoding/json"
	"net/http"

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
