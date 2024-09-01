package task

import (
	"time"
)

type TaskStatus string

const (
	StatusTodo TaskStatus = "todo"
	StatusDone TaskStatus = "done"
	StatusInProgress TaskStatus = "in progress"
)

type Task struct {
	ID			int				`json:"id"`
	Description string			`json:"description"`
	Status		TaskStatus		`json:"status"`
	CreatedAt   time.Time		`json:"created_at"`
	UpdatedAt   time.Time		`json:"updated_at"`
}

func AddTask(description string) error {
	// TODO
	return nil
}

func RemoveTask(id int) error {
	// TODO
	return nil
}

func UpdateTask(id int, status TaskStatus) error {
	// TODO
	return nil
}

func GetTasks() ([]Task, error) {
	
	return loadTasksFromJson()
}

