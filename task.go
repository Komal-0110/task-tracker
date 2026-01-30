package main

import "time"

type Status string

const (
	TODO       Status = "todo"
	InProgress Status = "inprogress"
	Done       Status = "done"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
