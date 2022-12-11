package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func (s *queue) Consume() ([]*QueueMessage, error) {
	return s.ConsumeWithContext(context.Background())
}

func (s *queue) ConsumeWithContext(ctx context.Context) ([]*QueueMessage, error) {
	queueInput := s.input
	ctx, cancel := context.WithTimeout(ctx, queueInput.Timeout)
	defer cancel()

	// Bate na fila sqs na AWS e retorna os objetos da fila. ReceiveMessageInput retorna 1 ou até 10 mensagens por vez
	receiveMessageOutput, err := s.client.ReceiveMessageWithContext(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:              aws.String(queueInput.URL),
		MaxNumberOfMessages:   aws.Int64(queueInput.MaxNumberOfMessages),
		WaitTimeSeconds:       aws.Int64(queueInput.WaitTimeSeconds),
		MessageAttributeNames: aws.StringSlice(queueInput.MessageAttributeNames),
	})

	if err != nil {
		return nil, err
	}

	queueMessage := make([]*QueueMessage, len(receiveMessageOutput.Messages))
	for i, msg := range receiveMessageOutput.Messages {
		queueMessage[i] = &QueueMessage{
			Body:          *msg.Body,
			receiptHandle: *msg.ReceiptHandle,
			queue:         s,
		}
	}
	return queueMessage, nil
}

func (s *queue) PollingIntervalSeconds() int64 {
	return s.input.PollingIntervalSeconds
}
