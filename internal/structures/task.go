package structures

import "time"

type Task struct {
	ID        int
	Message   string
	CreatedAt time.Time
	Status    StatusType
}

func NewTask(id int, msg string, statusType StatusType) *Task {
	return &Task{
		ID:        id,
		CreatedAt: time.Now(),
		Message:   msg,
		Status:    statusType,
	}
}
