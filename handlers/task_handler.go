package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/saloneepathan/task-manager-api/models"
	"github.com/saloneepathan/task-manager-api/storage"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.Tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range storage.Tasks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = uuid.New().String()
	task.Status = "pending"
	storage.Tasks = append(storage.Tasks, task)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range storage.Tasks {
		if item.ID == params["id"] {
			_ = json.NewDecoder(r.Body).Decode(&storage.Tasks[index])
			storage.Tasks[index].ID = item.ID
			json.NewEncoder(w).Encode(storage.Tasks[index])
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range storage.Tasks {
		if item.ID == params["id"] {
			storage.Tasks = append(storage.Tasks[:index], storage.Tasks[index+1:]...)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}
