package tasks

import (
	"errors"
	"sync"
	"time"
)

// Хранилище
var (
	tasks = make(map[string]*Task)
	mu    sync.RWMutex
)

// Добавить задачу
func AddTask(task *Task) {
	mu.Lock()
	tasks[task.ID] = task
	mu.Unlock()
}

// Получение задачи
func GetTask(id string) (*Task, error) {
	mu.RLock()
	defer mu.RUnlock()

	task, ok := tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return task, nil
}

// Обновить задачу
func UpdateTaskStatus(id string, status TaskStatus, startedAt, finishedAt *time.Time, errMsg string, result interface{}) error {
	mu.Lock()
	defer mu.Unlock()

	task, ok := tasks[id]
	if !ok {
		return errors.New("task not found")
	}

	task.Status = status
	if startedAt != nil {
		task.StartedAt = startedAt
	}
	if finishedAt != nil {
		task.FinishedAt = finishedAt
	}
	task.Error = errMsg
	if result != nil {
		task.Result = result
	}
	return nil
}

// Удалить задачу
func DeleteTask(id string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := tasks[id]; !ok {
		return errors.New("task not found")
	}
	delete(tasks, id)
	return nil
}
