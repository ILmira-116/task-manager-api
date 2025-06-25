package router

import (
	"task_manager_api/internal/handlers"
	"task_manager_api/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Логируем запросы
	r.Use(middleware.RequestLogger)

	// Создать задачу
	r.Post("/tasks", handlers.CreateTaskHandler)

	//Получить статус задачи
	r.Get("/tasks/{id}", handlers.GetTaskHandler)

	//Удалить задачу
	r.Delete("/tasks/{id}", handlers.DeleteTaskHandler)

	return r

}
