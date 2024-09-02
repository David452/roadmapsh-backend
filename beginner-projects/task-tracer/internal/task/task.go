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
	
	tasks, err := loadTasksFromJson()

	if err != nil {
		return fmt.Errorf("error loading tasks from JSON file: %w", err)
	}

	var newTasks []Task
	var idErr error = nil

	for _, t := range tasks {
		if t.ID == id {
			newTasks = append(tasks[:t.ID-1], tasks[t.ID:]...)
			idErr = nil
			break
		}
		idErr = fmt.Errorf("no task with ID %d found", id)
	}

	if idErr != nil {
		return idErr
	}

	err = saveTasksToJson(newTasks)

	return err
}

func UpdateTask(id int, newDescription string) error {

	tasks, err := loadTasksFromJson()

	if err != nil {
		return fmt.Errorf("error loading tasks from JSON file: %w", err)
	}
	var idErr error = nil

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()
			idErr = nil
			break
		}
		idErr = fmt.Errorf("no task with ID %d found", id)
	}

	if idErr != nil {
		return idErr
	}

	err = saveTasksToJson(tasks)

	return err
}

func GetTasks() ([]Task, error) {
	
	return loadTasksFromJson()
}

func MarkTask(id int, status TaskStatus) error {

	tasks, err := loadTasksFromJson()

	if err != nil {
		return fmt.Errorf("error loading tasks from JSON file: %w", err)
	}

	var idErr error = nil

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			idErr = nil
			break
		}
		idErr = fmt.Errorf("no task with ID %d found", id)
	}

	if idErr != nil {
		return idErr
	}

	err = saveTasksToJson(tasks)

	return err
}

