#aws-cli>=2.4.9

DYNAMO_TABLE_NAME=learn-memdb-table

list-tables:
	aws dynamodb list-tables --endpoint-url http://localhost:8000

create-table:
	aws dynamodb create-table \
		--table-name learn-memdb-table \
		--attribute-definitions \
			AttributeName=key,AttributeType=S \
		--key-schema \
			AttributeName=key,KeyType=HASH \
		--provisioned-throughput \
			ReadCapacityUnits=5,WriteCapacityUnits=5 \
		--table-class STANDARD \
		--provisioned-throughput \
			ReadCapacityUnits=5,WriteCapacityUnits=5 \
		--endpoint-url \
			http://localhost:8000