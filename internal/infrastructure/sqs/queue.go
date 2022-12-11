package sqs

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type queue struct {
	client *sqs.SQS
	input  Input
}

// New returns a new queue
func New(input Input) (Consumer, error) {
	i := newInput(input)
	s, erro := newSession(i)
	if erro != nil {
		return nil, erro
	}
	sqs := newQueue(s, i)
	return &sqs, nil
}

func newQueue(session *session.Session, input Input) queue {
	return queue{
		client: sqs.New(session),
		input:  input,
	}
}
