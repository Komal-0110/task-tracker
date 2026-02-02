package models

import (
	"time"

	"github.com/google/uuid"
)

type RunStatus string

const (
	RunQueued  RunStatus = "queued"
	RunRunning RunStatus = "running"
	RunSuccess RunStatus = "success"
	RunFailed  RunStatus = "failed"
)

type Run struct {
	Base
	TaskID       uuid.UUID `db:"task_id"`
	ScheduledFor time.Time `db:"scheduled_for"`
	StartedAt    time.Time `db:"started_at"`
	FinishedAt   time.Time `db:"finished_at"`
	Status       RunStatus `db:"status"`
	ErrorMessage string    `db:"error_message"`
	LLMModel     string    `db:"llm_model"`
	TokenUsage   []byte    `db:"token_usage"` // JSON
	CostEstimate float64   `db:"cost_estimate"`
}
