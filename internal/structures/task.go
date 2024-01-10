package structures

import "time"

type Task struct {
	ID        int
	Message   string
	CreatedAt time.Time
	DueDate   time.Time
	Status    string
}

func NewTask(id int, msg string, dueDate time.Time) *Task {
	return &Task{
		ID:        id,
		CreatedAt: time.Now(),
		Message:   msg,
		DueDate:   dueDate,
		Status:    "pending",
	}
}
