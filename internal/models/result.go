package models

import "github.com/google/uuid"

type Result struct {
	Base
	RunID uuid.UUID `db:"run_id"`
	Data  []byte    `db:"data"` // structured JSON result
}
