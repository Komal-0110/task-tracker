package main

import (
	"fmt"
	"time"
)

func addTask(description string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	now := time.Now()

	n := len(tasks)
	var newTaskId int
	if n == 0 {
		newTaskId = 1
	} else {
		newTaskId = n + 1
	}

	task := Task{
		Id:          newTaskId,
		Description: description,
		Status:      TODO,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, task)

	return saveTasks(tasks)
}

func updateStatus(task Task) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].Id == task.Id {
			tasks[i] = task
			tasks[i].UpdatedAt = time.Now()
			return saveTasks(tasks)
		}
	}

	return fmt.Errorf("task not found")
}

func listTasks(status *Status) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if status == nil || task.Status == *status {
			fmt.Printf("[%s] %s (%d)\n", task.Status, task.Description, task.Id)
		}
	}

	return nil
}

func updateTaskStatus(id int, status Status) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for i, t := range tasks {
		if t.Id == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()

			if err := saveTasks(tasks); err != nil {
				fmt.Printf("failed to update task: %v\n", err)
			}
			return
		}
	}

	fmt.Println("task not found")
}
