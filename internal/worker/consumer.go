package worker

import (
	"learn-memdb/internal/domain/appcontext"
)

// Consumer is the interface that receive message from SQS
type Consumer interface {
	URL() string
	Handler(appcontext.Context, Input, SqsEntryEntity) error
	PollingIntervalSeconds() int64
}
