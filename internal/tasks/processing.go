package tasks

import (
	"math/rand"
	"time"
)

// Обрабатываем задачу в отдельной горутине
func StartTaskProcessing(task *Task) {
	go func(t *Task) {
		// Меняем статус на RUNNING
		now := time.Now()
		UpdateTaskStatus(t.ID, StatusRunning, &now, nil, "", nil)

		// Имитация долгой работы
		time.Sleep(3 * time.Minute)

		// Завершение обработки
		finished := time.Now()
		if rand.Intn(100) < 70 { // 70% успеха
			UpdateTaskStatus(t.ID, StatusSuccess, nil, &finished, "", "Результат успешный")
		} else {
			UpdateTaskStatus(t.ID, StatusFailed, nil, &finished, "Произошла ошибка", nil)
		}
	}(task)
}
