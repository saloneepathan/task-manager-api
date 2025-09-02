package main

import (
	"log"
	"net/http"

	"github.com/saloneepathan/task-manager-api/routes"
)

func main() {
	router := routes.RegisterRoutes()

	log.Println("🚀 Task Manager API running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
