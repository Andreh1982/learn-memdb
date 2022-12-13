LOCALSTACK=http://localhost:4566
QUEUE=sre-dev
QUEUE_URL=${LOCALSTACK}/000000000000/${QUEUE}

queue-send-message:
	aws --endpoint-url=${LOCALSTACK} sqs send-message --queue-url ${QUEUE_URL} --message-body '{"createdAt": "2022-12-11T00:57:58.085837193-03:00", "updatedAt": "0001-01-01T00:00:00Z", "entryId": "9779d276-9999-9999-9999-445e44319e1d", "name": "letter-to-memphins"}'

queue-create:
	aws --endpoint-url=${LOCALSTACK} sqs create-queue --queue-name ${QUEUE}

queue-list:
	aws --endpoint-url=${LOCALSTACK} sqs list-queues

queue-receive-messages:
	aws --endpoint-url=${LOCALSTACK} sqs receive-message --queue-url ${QUEUE_URL} --max-number-of-messages 10
