package worker

import (
	"learn-memdb/internal/domain/appcontext"
	"learn-memdb/internal/infrastructure/environment"
)

type consumerWhoDeletes struct{}

func (c *consumerWhoDeletes) URL() string {
	return environment.GetInstance().AWS_SQS_URL_TO_DELETE
}

func (c *consumerWhoDeletes) Handler(ctx appcontext.Context, input Input, message SqsEntryEntity) error {
	err := input.LearnMemdbUseCases.Delete(ctx, message.UUID)
	return err
}

func (c *consumerWhoDeletes) PollingIntervalSeconds() int64 {
	return environment.GetInstance().INTERVAL_GET_KEYS_TO_DELETE
}
