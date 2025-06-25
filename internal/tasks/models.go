package tasks

import "time"

type TaskStatus string

const (
	StatusPending TaskStatus = "PENDING"
	StatusRunning TaskStatus = "RUNNING"
	StatusSuccess TaskStatus = "SUCCESS"
	StatusFailed  TaskStatus = "FAILED"
)

type Task struct {
	ID         string     `json:"id"`
	Status     TaskStatus `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	StartedAt  *time.Time `json:"started_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	Error      string     `json:"error,omitempty"`
	Result     any        `json:"result,omitempty"`
}

func (t *Task) ProcessingDuration() *time.Duration {
	if t.StartedAt == nil {
		return nil
	}

	var end time.Time
	if t.FinishedAt != nil {
		end = *t.FinishedAt
	} else {
		end = time.Now()
	}

	duration := end.Sub(*t.StartedAt)
	return &duration
}
