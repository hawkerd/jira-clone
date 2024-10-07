package handlers

import (
	"fmt"
	"net/http"

	"github.com/hawkerd/jira-clone/database"
	"github.com/hawkerd/jira-clone/models"
)

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// ensure name was provided
		name := r.FormValue("name")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Project name is required")
			return
		}

		// create new project
		newProject := models.Project{
			Name: name,
		}

		// save the project to the database
		database.DB.Create(&newProject)

		// respond with details about the created project
		fmt.Fprintf(w, "Created Project: %d, Name: %s\n", newProject.ID, newProject.Name)

	case http.MethodGet:
		var projects []models.Project

		// retrieve all projects from the database
		database.DB.Find(&projects)

		// respond with information about all projects
		if len(projects) == 0 {
			fmt.Fprintln(w, "No projects available")
			return
		} else {
			for _, project := range projects {
				fmt.Fprintf(w, "Project ID: %d, Name: %s\n", project.ID, project.Name)
			}

		}
	}
}
