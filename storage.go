package main

import (
	"encoding/json"
	"os"
)

const fileName = "tasks.json"

func loadTasks() ([]Task, error) {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}

		return []Task{}, err
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		return []Task{}, err
	}

	var task []Task
	if err := json.Unmarshal(content, &task); err != nil {
		return []Task{}, err
	}

	return task, nil
}
