package worker

import (
	"encoding/json"
	"fmt"

	"learn-memdb/internal/domain/appcontext"
	"learn-memdb/internal/infrastructure/environment"
	"learn-memdb/internal/infrastructure/sqs"
	"learn-memdb/internal/infrastructure/sqs/polling"

	"go.uber.org/zap"
)

func createPolling(ctx appcontext.Context, input Input, consumer Consumer) {
	logger := input.Logger
	sqsInput := sqs.Input{
		URL:                    consumer.URL(),
		PollingIntervalSeconds: consumer.PollingIntervalSeconds(),
		AwsEndpoint:            environment.GetInstance().SQS_AWS_ENDPOINT,
	}
	logger.Info("Starting polling", zap.String("URL", sqsInput.URL), zap.Int64("PollingIntervalSeconds", sqsInput.PollingIntervalSeconds))
	for event := range polling.Consume(sqsInput) {
		err := event.Args[1]
		if err != nil {
			logger.Error("Error fetch message", zap.Any("Error", err))
			continue
		}
		qm := event.Args[0].(*sqs.QueueMessage)

		var sqsEntryEntity SqsEntryEntity
		if err := json.Unmarshal([]byte(qm.Body), &sqsEntryEntity); err != nil {
			logger.Error("Error parse message", zap.Any("Error", err))
			continue
		}
		err = consumer.Handler(ctx, input, sqsEntryEntity)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}
		_, err = qm.DeleteMessage()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			logger.Error("Error on remove message", zap.Any("Error", err))
			continue
		}
		fmt.Printf("\"Mensagem removida da fila\": %v\n", "Mensagem removida da fila")
	}
}
