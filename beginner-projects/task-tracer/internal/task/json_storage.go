package task

import (
	"encoding/json"
	"os"
	"fmt"
)

// The file where tasks are stored
const jsonFile = "tasks.json"

func saveTasksToJson(tasks []Task) error {

	jsonData, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		return fmt.Errorf("error converting structs to JSON format: %w", err)
	}

	err = os.WriteFile(jsonFile, jsonData, 0644)

	if err != nil {
		return fmt.Errorf("error writing JSON data to file: %w", err)
	}

	return nil
}

func loadTasksFromJson() ([]Task, error) {

	// Check if the file exists
	if _, err := os.Stat(jsonFile); os.IsNotExist(err) {

		file, err := os.Create(jsonFile)

		if err != nil {
			return nil, fmt.Errorf("error creating JSON file: %w", err)
		}
		defer file.Close()

		// Initialize the file with an empty JSON array
		_, err = file.Write([]byte("[]"))

		if err != nil {
			return nil, fmt.Errorf("error writing empty JSON array to file: %w", err)
		}

		return []Task{}, nil
	}
	
	// Open the file
	file, err := os.Open(jsonFile)

	if err != nil {
		return nil, fmt.Errorf("error opening JSON file: %w", err)
	}
	defer file.Close()

	// Decode json
	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON data: %w", err)
	}
	
	return tasks, nil
}
