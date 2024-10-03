package main

import (
	"fmt"
	"net/http"
)

func homeHandler(write http.ResponseWriter, read *http.Request) {
	fmt.Fprintln(write, "Welcome to your Jira clone!")
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Starting server on port :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
