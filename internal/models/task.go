package models

import "time"

type TaskStatus string

const (
	TaskEnabled  TaskStatus = "enabled"
	TaskDisabled TaskStatus = "disabled"
)

type Task struct {
	Base
	Name             string     `db:"name"`
	Prompt           string     `db:"prompt"`
	CronExpression   string     `db:"cron_expression"`
	Timezone         string     `db:"timezone"`
	WebSearchEnabled bool       `db:"web_search_enabled"`
	Status           TaskStatus `db:"status"`
	NextRunAt        time.Time  `db:"next_run_at"`
}
