package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_loadTasks(t *testing.T) {
	tests := []struct {
		name      string
		fileData  string
		wantTasks []Task
		wantErr   bool
	}{
		{
			name:      "invalid json :NEG",
			fileData:  "invalid json",
			wantTasks: []Task{},
			wantErr:   true,
		},
		{
			name:      "file does not exist :POS",
			fileData:  "",
			wantTasks: []Task{},
			wantErr:   false,
		},
		{
			name: "get tasks :POS",
			fileData: `[
				{
					"id": 1,
					"description": "this is dummy task",
					"status": "inprogress",
					"created_at": "2024-01-01T10:00:00Z",
					"updated_at": "2024-01-01T10:00:00Z"
				},
				{
					"id": 2,
					"description": "this is dummy task2",
					"status": "done",
					"created_at": "2024-01-01T11:00:00Z",
					"updated_at": "2024-01-01T12:00:00Z"
				}
			]
			`,
			wantTasks: []Task{
				{
					Id:          1,
					Description: "this is dummy task",
					Status:      InProgress,
					CreatedAt:   time.Date(2024, 01, 1, 10, 00, 00, 00, time.UTC),
					UpdatedAt:   time.Date(2024, 01, 1, 10, 00, 00, 00, time.UTC),
				},
				{
					Id:          2,
					Description: "this is dummy task2",
					Status:      Done,
					CreatedAt:   time.Date(2024, 01, 1, 11, 00, 00, 00, time.UTC),
					UpdatedAt:   time.Date(2024, 01, 1, 12, 00, 00, 00, time.UTC),
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()

			wd, err := os.Getwd()
			if err != nil {
				t.Fatalf("failed to get wd: %v", err)
			}
			defer os.Chdir(wd)

			if err := os.Chdir(tempDir); err != nil {
				t.Fatalf("failed to chdir: %v", err)
			}

			if tt.fileData != "" {
				fileName := filepath.Join(tempDir, "tasks.json")
				if err := os.WriteFile(fileName, []byte(tt.fileData), 0644); err != nil {
					t.Fatalf("failed to write into file: %v", err)
				}
			}

			gotTasks, gotErr := loadTasks()

			if tt.wantErr {
				assert.Error(t, gotErr)
			} else {
				assert.NoError(t, gotErr)
			}

			assert.Equal(t, tt.wantTasks, gotTasks, "expect tasks to match")
		})
	}
}

func Test_saveTasks(t *testing.T) {
	now := time.Now().UTC()

	tests := []struct {
		name  string
		tasks []Task
	}{
		{
			name:  "add empty list of task :POS",
			tasks: []Task{},
		},
		{
			name: "add list of tasks :POS",
			tasks: []Task{
				{
					Id:          1,
					Description: "this is dummy task",
					Status:      TODO,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()

			wd, err := os.Getwd()
			if err != nil {
				t.Fatalf("failed to get wd: %v", err)
			}
			defer os.Chdir(wd)

			if err := os.Chdir(tempDir); err != nil {
				t.Fatalf("failed to chdir: %v", err)
			}

			gotErr := saveTasks(tt.tasks)

			assert.NoError(t, gotErr)

			content, err := os.ReadFile("tasks.json")
			if err != nil {
				t.Error("failed to get content of tasks.json", err)
			}

			var got []Task
			if err := json.Unmarshal(content, &got); err != nil {
				t.Error("failed to unmarshal content", err)
			}

			assert.Equal(t, tt.tasks, got, "expect task list to match")
		})
	}
}
