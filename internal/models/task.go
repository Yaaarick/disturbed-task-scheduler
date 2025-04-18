package models

import (
	"encoding/json"
	"time"
)

// TaskStatus represents the status of a task
type TaskStatus string

const (
	TaskStatusPending  TaskStatus = "pending"
	TaskStatusRunning  TaskStatus = "running"
	TaskStatusComplete TaskStatus = "complete"
	TaskStatusFailed   TaskStatus = "failed"
	TaskStatusRetrying TaskStatus = "retrying"
)

// TaskPriority represents the priority of a task
type TaskPriority int

const (
	PriorityLow    TaskPriority = 1
	PriorityNormal TaskPriority = 2
	PriorityHigh   TaskPriority = 3
)

// Task represents a task entity
type Task struct {
	ID          string          `json:"id" db:"id"`
	QueueName   string          `json:"queue_name" db:"queue_name"`
	TaskType    string          `json:"task_type" db:"task_type"`
	PayLoad     json.RawMessage `json:"pay_load" db:"pay_load"`
	Priority    TaskPriority    `json:"priority" db:"priority"`
	Status      TaskStatus      `json:"status" db:"status"`
	RetryCount  int             `json:"retry_count" db:"retry_count"`
	MaxRetries  int             `json:"max_retries" db:"max_retries"`
	LastError   string          `json:"last_error" db:"last_error"`
	CreatedAt   time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" db:"updated_at"`
	NextRunAt   time.Time       `json:"next_run_at" db:"next_run_at"`
	CompletedAt *time.Time      `json:"completed_at" db:"completed_at"`
	WorkerID    *string         `json:"worker_id" db:"worker_id"`
}
