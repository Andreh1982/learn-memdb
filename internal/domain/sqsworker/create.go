package worker

import (
	"learn-memdb/internal/domain/appcontext"
	"learn-memdb/internal/domain/learnmemdb"
	"learn-memdb/internal/infrastructure/environment"
)

type consumerWhoCreates struct{}

func (c *consumerWhoCreates) URL() string {
	return environment.GetInstance().AWS_SQS_URL_QUEUE
}

func (c *consumerWhoCreates) Handler(ctx appcontext.Context, input Input, message SqsEntryEntity) error {
	entryEntity := &learnmemdb.EntryEntity{
		EntryID:   message.EntryID,
		IsFromSqs: true,
	}

	_, err := input.LearnMemdbUseCases.Create(ctx, *entryEntity)
	return err
}

func (c *consumerWhoCreates) PollingIntervalSeconds() int64 {
	return environment.GetInstance().INTERVAL_GET_KEYS
}
