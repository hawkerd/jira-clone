package models

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
	ProjectID   int
}
