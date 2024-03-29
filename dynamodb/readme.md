# ddb docker image 參數操作方法
https://docs.aws.amazon.com/zh_tw/amazondynamodb/latest/developerguide/DynamoDBLocal.UsageNotes.html

# How to use
1. `docker-compose up -d`
2. open the dynamodbAdmin window
3. `docker exec -it dynamodbClient go run main.go`
4. Observe the dynamodbAdmin:test

# 關於 localhost_dynamodb 
## 載入自定義資料庫
- 參考 docker-compose.yml
- 像 mysql 一樣, 可以預載
- 但會依據操作, 資料內容會變動
  - 建立用 script 塞入預設資料
  - 或者不更動預設資料庫, 用 cp 方式, 複製一份新的操作
## localhost 跟 remote 的 ddb 差異
- https://docs.aws.amazon.com/zh_tw/amazondynamodb/latest/developerguide/DynamoDBLocal.UsageNotes.html



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


# Scan 比較運算子與函數參考

> https://docs.aws.amazon.com/zh_tw/amazondynamodb/latest/developerguide/Expressions.OperatorsAndFunctions.html

- 可在 scan 中，增加以下多種不同的條件，透過 `FilterExpression` 來設定
    - attribute_not_exists
    - attribute_exists
    - begins_with
    - contains
    - 或者其他運算子: <=, >=, >, < ....

# 使用 `dynamodbattribute`, 將內容轉換成定義好的 struct format

```
var objs []OBJS
err = dynamodbattribute.UnmarshalListOfMaps(PutItemOutput.Items, &objs)
if err != nil {
    panic(err)
}
```