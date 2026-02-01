package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task [add|list|done|progress]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("description required")
			return
		}
		desc := strings.Join(os.Args[2:], " ")
		err := addTask(desc)
		if err != nil {
			fmt.Printf("failed to add task:%s", err)
		}

	case "done":
		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("failed to convert id string to int:%s", err)
			return
		}
		updateTaskStatus(id, Done)

	case "progress":
		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("failed to convert id string to int:%s", err)
			return
		}
		updateTaskStatus(id, InProgress)

	case "list":
		listTasks(nil)

	case "list-done":
		s := Done
		listTasks(&s)

	case "list-pending":
		s := TODO
		listTasks(&s)

	case "list-progress":
		s := InProgress
		listTasks(&s)

	default:
		fmt.Printf("wrong args:%s", command)
	}

}
