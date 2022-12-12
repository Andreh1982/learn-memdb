LOCALSTACK=http://localhost:4566
QUEUE=sre-dev
QUEUE_URL=${LOCALSTACK}/000000000000/${QUEUE}

queue-send-message:
	aws --endpoint-url=${LOCALSTACK} sqs send-message --queue-url ${QUEUE_URL} --message-body '{     "createdAt": "milky-way",     "updatedAt": "00000000-0000-0000-0000-000000000000",     "entryId": "app-00000000-0000-0000-0000-000000000000.ops.olist.io",     "name": "letter-to-memphins" }'

queue-create:
	aws --endpoint-url=${LOCALSTACK} sqs create-queue --queue-name ${QUEUE}

queue-list:
	aws --endpoint-url=${LOCALSTACK} sqs list-queues

queue-receive-messages:
	aws --endpoint-url=${LOCALSTACK} sqs receive-message --queue-url ${QUEUE_URL} --max-number-of-messages 10
