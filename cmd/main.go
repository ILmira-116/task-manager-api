package main

import (
	"net/http"
	"task_manager_api/internal/logger"
	"task_manager_api/internal/router"
)

func main() {
	// Инициализируем логгер с debug = true (для примера)
	logger.Init(true)

	logger.Log.Info("Starting server on :8080")

	r := router.NewRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Log.Fatalf("Server failed: %v", err)
	}

}
