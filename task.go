package main

import "time"

type Status string

const (
	TODO       Status = "todo"
	InProgress Status = "inprogress"
	Done       Status = "done"
)

type Task struct {
	Id          string
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
