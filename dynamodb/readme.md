# How to use
1. `docker-compose up -d`
2. open the dynamodbAdmin window
3. `docker exec -it dynamodbClient go run main.go`
4. Observe the dynamodbAdmin:test

# Others
## dynamodbAdmin : 
- `http://localhost:8001`

## Scan
> https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.Scan

### ScanInput

#### ExpressionAttributeValues
- 定義 value 屬性 : https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#AttributeValue


## awscli : dynamodb
- https://docs.aws.amazon.com/cli/latest/reference/dynamodb/index.html

### create table

```
awscli --endpoint-url http://dynamodb:8000 dynamodb create-table \
--table-name test \
--key-schema AttributeName=id,KeyType=HASH \
--attribute-definitions AttributeName=id,AttributeType=S \
--provisioned-throughput ReadCapacityUnits=3,WriteCapacityUnits=3
```

### insert new datas

```
awscli --endpoint-url http://dynamodb:8000 dynamodb batch-write-item --request-items file://default_datas.json
```

- about `default_datas.json` : The table name need to identify. Present name is `test`