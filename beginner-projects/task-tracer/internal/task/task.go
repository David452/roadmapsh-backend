package task

import (
	"time"
	"fmt"
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
	tasks, _ := loadTasksFromJson()

	newTask := Task {
		ID: len(tasks) + 1,
		Description: description,
		Status: StatusTodo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)

	err := saveTasksToJson(tasks)

	if err != nil {
		return fmt.Errorf("error saving tasks to JSON file: %w", err)
	}
	
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

