package sqs

import (
	"context"
)

// Consumer is the interface that receive message from SQS
type Consumer interface {
	Consume() ([]*QueueMessage, error)
	ConsumeWithContext(ctx context.Context) ([]*QueueMessage, error)
	PollingIntervalSeconds() int64
}
