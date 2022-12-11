LOCALSTACK=http://localhost:4566
QUEUE_TO_CREATE=sre-dev-to-create
QUEUE_URL_TO_CREATE=${LOCALSTACK}/000000000000/${QUEUE_TO_CREATE}
QUEUE_TO_DELETE=sre-dev-to-delete
QUEUE_URL_TO_DELETE=${LOCALSTACK}/000000000000/${QUEUE_TO_DELETE}

queue-create-message-to-create:  ## Send a message to the localstack queue
	aws --endpoint-url=${LOCALSTACK} sqs send-message --queue-url ${QUEUE_URL_TO_CREATE} --message-body '{     "from": "OPS_WEBAPP_CREATED",     "type": "webapp",     "uuid": "00000000-0000-0000-0000-000000000000",     "name": "hello-world",     "namespace_uuid": "00000000-0000-0000-0000-000000000000",     "namespace_name": "earth",     "cluster_alias": "milky-way",     "instance_uuid": "00000000-0000-0000-0000-000000000000",     "address": "app-00000000-0000-0000-0000-000000000000.ops.olist.io",     "gitrepo": "git@git.ops.olist.io:00000000-0000-0000-0000-000000000000.git" }'

queue-create-message-to-delete:  ## Send a message to the localstack queue
	aws --endpoint-url=${LOCALSTACK} sqs send-message --queue-url ${QUEUE_URL_TO_DELETE} --message-body '{     "from": "OPS_WEBAPP_DELETED",     "type": "webapp",     "uuid": "00000000-0000-0000-0000-000000000000",     "name": "hello-world",     "namespace_uuid": "00000000-0000-0000-0000-000000000000",     "namespace_name": "earth",     "instance_uuid": "00000000-0000-0000-0000-000000000000",     "cluster_alias": "milky-way" }'

queue-create:  ## Create the localstack queue
	aws --endpoint-url=${LOCALSTACK} sqs create-queue --queue-name ${QUEUE_TO_CREATE}
	aws --endpoint-url=${LOCALSTACK} sqs create-queue --queue-name ${QUEUE_TO_DELETE}

queue-list:  ## List the localstack queue
	aws --endpoint-url=${LOCALSTACK} sqs list

queue-receive:  ## Create the localstack queue
	aws --endpoint-url=${LOCALSTACK} sqs receive-message --queue-url ${QUEUE_URL_TO_CREATE} --max-number-of-messages 10
