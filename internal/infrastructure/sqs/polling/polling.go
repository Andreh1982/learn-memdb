package polling

import (
	"context"
	"fmt"
	"time"

	"learn-memdb/internal/infrastructure/sqs"

	"github.com/olebedev/emitter"
)

// Consume is a wrapper for Consume[T] that allows to pass a context
func Consume(input sqs.Input) <-chan emitter.Event {
	return ConsumeWithContext(context.Background(), input)
}

// ConsumeWithContext is a consumer that will consume messages from a SQS queue
// and emit them as events
// It will also handle the SQS errors and emit them as events
// It will also handle the SQS messages that are not valid JSON and emit them as events
// BreakChange: https://github.com/swaggo/swag/issues/1170
func ConsumeWithContext(ctx context.Context, input sqs.Input) <-chan emitter.Event {
	e := &emitter.Emitter{}
	consumer, err := sqs.New(input)
	const event = "update"
	if err != nil {
		fmt.Printf("err: %v\n", err)
		<-e.Emit(event, nil, err)
	}

	go func() {
		for {
			time.Sleep(time.Duration(consumer.PollingIntervalSeconds()) * time.Second)
			queueMessageList, err := consumer.Consume()

			if err != nil {
				fmt.Printf("err: %v\n", err)
				<-e.Emit(event, nil, err)
				continue
			}

			for _, queueMessage := range queueMessageList {
				<-e.Emit(event, queueMessage, nil)
			}
		}
	}()

	return e.On(event)
}
