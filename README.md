# disturbed-task-scheduler
A lightweight Celery/Kafka-like system for distributed task processing with queues, retries, and priorities.  Go NATS PostgreSQL Prometheus

Features:
- Distributed task processing: Workers pull tasks via NATS (competing consumers pattern).

- At-least-once delivery: Tasks are retried on failures (with exponential backoff).

- Priority queues: High/medium/low priority task routing.

- Observability: Prometheus metrics (task latency, success rates) + Grafana dashboards.

- Graceful shutdown: In-flight tasks complete before worker termination.

- Dead Letter Queue (DLQ): Isolate poison pills for debugging.
