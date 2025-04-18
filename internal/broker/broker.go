package broker

import (
	"context"
	"disturbed-task-scheduler/internal/models"
	"time"
)

// Broker interface defines the methods for a message broker
type Broker interface {
	// PublishTask publishes a task to the broker
	PublishTask(ctx context.Context, task *models.Task) error

	// ConsumeTask consumes a task from the broker
	ConsumeTask(ctx context.Context, queueName string) (*models.Task, error)

	// AckTask acknowledges a task as completed
	AckTask(ctx context.Context, taskID string) error

	// NackTask nacks a task as failed
	NackTask(ctx context.Context, taskID string, err error) error

	// RetryTask retries a task
	RetryTask(ctx context.Context, taskID string, nextRetryAt time.Time) error
}
