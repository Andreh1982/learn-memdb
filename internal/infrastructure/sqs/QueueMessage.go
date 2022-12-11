package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// QueueMessage is the message from SQS
type QueueMessage struct {
	Body          string
	receiptHandle string
	queue         *queue
	Data          interface{}
}

// DeleteMessage deletes the message from SQS
func (qm *QueueMessage) DeleteMessage() (string, error) {
	return qm.DeleteMessageWithContext(context.Background())
}

// DeleteMessageWithContext deletes the message from SQS
func (qm *QueueMessage) DeleteMessageWithContext(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, qm.queue.input.Timeout)
	defer cancel()

	_, err := qm.queue.client.DeleteMessageWithContext(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(qm.queue.input.URL),
		ReceiptHandle: aws.String(qm.receiptHandle),
	})

	if err != nil {
		return "", err
	}

	return qm.receiptHandle, nil
}
