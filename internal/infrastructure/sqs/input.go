package sqs

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// Input is the input for the SQS consumer
type Input struct {
	AwsRegion              string
	AwsProfile             string
	AwsEndpoint            string
	MaxNumberOfMessages    int64
	WaitTimeSeconds        int64
	MessageAttributeNames  []string
	Timeout                time.Duration
	URL                    string
	PollingIntervalSeconds int64
}

func newInput(input Input) Input {

	i := Input{
		AwsRegion:              input.AwsRegion,
		AwsProfile:             input.AwsProfile,
		AwsEndpoint:            input.AwsEndpoint,
		Timeout:                input.Timeout,
		MaxNumberOfMessages:    input.MaxNumberOfMessages,
		WaitTimeSeconds:        input.WaitTimeSeconds,
		MessageAttributeNames:  input.MessageAttributeNames,
		URL:                    input.URL,
		PollingIntervalSeconds: input.PollingIntervalSeconds,
	}

	if i.AwsRegion == "" {
		i.AwsRegion = getenv("AWS_REGION", "us-east-1")
	}

	if i.AwsProfile == "" {
		i.AwsProfile = getenv("AWS_PROFILE", "default")
	}

	if i.AwsEndpoint == "" {
		i.AwsEndpoint = getenv("AWS_ENDPOINT", "http://localhost:4566")
	}

	if i.Timeout == 0 {
		timeout := getenvDuration("SQS_TIMEOUT", "60")
		i.Timeout = time.Second * timeout
	}

	if i.MaxNumberOfMessages == 0 {
		i.MaxNumberOfMessages = getenvInt64("SQS_MAX_NUMBER_OF_MESSAGES", "10")
	}

	if i.WaitTimeSeconds == 0 {
		i.WaitTimeSeconds = getenvInt64("SQS_WAIT_TIME_SECONDS", "5")
	}

	if i.MessageAttributeNames == nil {
		messageAttributeNames := getenv("SQS_MESSAGE_ATTRIBUTE_NAMES", "All")
		i.MessageAttributeNames = strings.Split(messageAttributeNames, ",")
	}

	if i.PollingIntervalSeconds == 0 {
		i.PollingIntervalSeconds = getenvInt64("SQS_POLLING_INTERVAL_SECONDS", "30")
	}

	return i
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getenvDuration(key, fallback string) time.Duration {
	value := getenvInt(key, fallback)
	valueFloat := time.Duration(value)
	return valueFloat
}

func getenvInt64(key, fallback string) int64 {
	value := getenv(key, fallback)
	valueFloat, _ := strconv.ParseInt(value, 10, 64)
	return valueFloat
}

func getenvInt(key, fallback string) int {
	value := getenv(key, fallback)
	valueFloat, _ := strconv.Atoi(value)
	return valueFloat
}
