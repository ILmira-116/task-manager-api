package handlers

import (
	"encoding/json"
	"net/http"
	"task_manager_api/internal/logger"
	"task_manager_api/internal/tasks"
	"time"

	"github.com/go-chi/chi/v5"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := tasks.GenerateSimpleID()
	createdAt := time.Now()

	task := &tasks.Task{
		ID:        id,
		Status:    tasks.StatusPending,
		CreatedAt: createdAt,
	}

	tasks.AddTask(task)
	tasks.StartTaskProcessing(task)

	logger.Log.Infof("Task created: ID=%s", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, err := tasks.GetTask(id)
	if err != nil {
		logger.Log.Warnf("Task not found: ID=%s", id)
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	logger.Log.Infof("Task fetched: ID=%s", id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := tasks.DeleteTask(id)
	if err != nil {
		logger.Log.Warnf("Task not found for deletion: ID=%s", id)
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	logger.Log.Infof("Task deleted: ID=%s", id)
	w.WriteHeader(http.StatusNoContent)
}
