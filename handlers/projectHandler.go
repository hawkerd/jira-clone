package handlers

import (
	"fmt"
	"net/http"

	"github.com/hawkerd/jira-clone/models"
)

var projects []models.Project
var projectID int = 0

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

		// create neew project and append it to the list of projects
		projectID++
		newProject := models.Project{
			ID:   projectID,
			Name: name,
		}
		projects = append(projects, newProject)

		fmt.Fprintf(w, "Created Project: %d, Name: %s\n", newProject.ID, newProject.Name)

	case http.MethodGet:
		if len(projects) == 0 {
			fmt.Fprintln(w, "No projects available")
			return
		}

		for _, project := range projects {
			fmt.Fprintf(w, "Project ID: %d, Name: %s\n", project.ID, project.Name)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Only GET and POST methods allowed")
	}

}
