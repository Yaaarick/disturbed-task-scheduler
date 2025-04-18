package storage

import (
	"context"
	"disturbed-task-scheduler/internal/models"
	"time"
)

type Storage interface {
	// Task operations
	CreateTask(ctx context.Context, task *models.Task) error
	GetTask(ctx context.Context, taskID string) (*models.Task, error)
	UpdateTaskStatus(ctx context.Context, taskID string, status models.TaskStatus) error
	ListTasks(ctx context.Context, filter TaskFilter) ([]models.Task, error)

	// Worker operations
	RegisterWorker(ctx context.Context, workerID string, hostname string) error
	UpdateWorkerHeartbeat(ctx context.Context, workerID string) error
	ListActiveWorkers(ctx context.Context) ([]Worker, error)

	// Task history operations
	RecordTaskExecution(ctx context.Context, history *TaskHistory) error
	GetTaskHistory(ctx context.Context, taskID string) ([]TaskHistory, error)
}

// TaskFilter represents the filter options for listing tasks
type TaskFilter struct {
	Statuses   []models.TaskStatus
	QueueNames []string
	FromDate   time.Time
	ToDate     time.Time
	Limit      int
	Offset     int
}

// Worker represents a worker entity
type Worker struct {
	ID            string
	Hostname      string
	LastHeartbeat time.Time
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// TaskHistory represents the history of a task execution
type TaskHistory struct {
	ID           string
	TaskID       string
	WorkerID     string
	Status       models.TaskStatus
	ErrorMessage string
	StartedAt    time.Time
	FinishedAt   time.Time
}
