package models

import "github.com/google/uuid"

type WebSearchSnapshot struct {
	Base
	RunID   uuid.UUID `db:"run_id"`
	Query   string    `db:"query"`
	Results []byte    `db:"results"` // JSON
}
