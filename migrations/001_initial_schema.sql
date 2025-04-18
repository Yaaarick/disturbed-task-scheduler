-- Create a type for task status
CREATE TYPE task_status AS ENUM (
    'pending',
    'running',
    'completed',
    'failed',
    'retrying'
);
-- Create a table for tasks
CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    queue_name VARCHAR(255) NOT NULL,
    task_type VARCHAR(255) NOT NULL,
    payload JSONB NOT NULL,
    priority SMALLINT NOT NULL DEFAULT 2,
    status task_status NOT NULL DEFAULT 'pending',
    retry_count INTEGER NOT NULL DEFAULT 0,
    max_retries INTEGER NOT NULL DEFAULT 3,
    last_error TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    next_run_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMP WITH TIME ZONE,
    worker_id VARCHAR(255),
    
    -- Indexes
    INDEX idx_tasks_status (status),
    INDEX idx_tasks_queue_priority (queue_name, priority),
    INDEX idx_tasks_next_run (next_run_at)
);

-- Create a table for workers
CREATE TABLE workers (
    id VARCHAR(255) PRIMARY KEY,
    hostname VARCHAR(255) NOT NULL,
    last_heartbeat TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Create a table for task history
CREATE TABLE task_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    task_id UUID NOT NULL REFERENCES tasks(id),
    worker_id VARCHAR(255) NOT NULL,
    status task_status NOT NULL,
    error_message TEXT,
    started_at TIMESTAMP WITH TIME ZONE NOT NULL,
    finished_at TIMESTAMP WITH TIME ZONE NOT NULL,
    
    INDEX idx_task_history_task_id (task_id)
);