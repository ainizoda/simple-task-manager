package task

import "time"

type TaskStatus string

const (
	Done       TaskStatus = "done"
	InProgress TaskStatus = "in progress"
	New        TaskStatus = "new"
)

type Task struct {
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
	Status      TaskStatus
}
