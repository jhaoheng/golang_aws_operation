# 目的
> https://docs.aws.amazon.com/zh_tw/serverlessrepo/latest/devguide/serverlessrepo-quick-start.html

1. 在本地端建立 Lambda func 與 apiGateway, 並且測試
2. 測試, 模擬 event : cus_event & s3Event
3. 部署到雲端
4. 在雲端測試

# doc
- SAM : https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md
- golang event 
    - doc : https://godoc.org/github.com/aws/aws-lambda-go/events
    - example : https://github.com/aws/aws-lambda-go/tree/master/events


# use sam to test
1. `make build`
2. 建立 : `sam local start-api --region us-east-1`
3. 測試 : 
    - `curl -X GET http://localhost:3000/hello/get`
    - `curl -X DELETE http://localhost:3000/hello/del`
    - `curl -X POST http://localhost:3000/hello/post -d '{"name":"hello this is your name"}'`
4. `make clean`


# use sam to invoke lambda
> `make build`

## invoke api
- `sam local invoke "apiget"`

## invoke cus-event
- `echo '{"message": "Hey, are you there?" }' | sam local invoke "cus_event" --event -`

## invoke s3 event
- `sam local generate-event s3 put --bucket my-bucket-20200508 --key hello | sam local invoke "s3Event" --event -`


# use sam to deploy
## 預先準備
- 部署流程
    - `https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-template-publishing-applications.html`
- metadata 內容範例
    - 必須手動建立, 必要流程, 否則 `sam public` 會出現錯誤
    - `https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-template-publishing-applications-metadata-properties.html`

``` # Metadata 必須標籤
Metadata:
  AWS::ServerlessRepo::Application:
    Name: my-app
    Description: hello world
    Author: user1
    SemanticVersion: 0.0.1
```

## 上傳 SAM 到 SAR, 並且從 SAR 中部署資源
1. 建立 packaged.yaml, 指定要儲存範本的 s3 bucket
    - 需要設定 s3 bucket policy : https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-template-publishing-applications.html
    - `sam package --template-file template.yaml --output-template-file packaged.yml --s3-bucket {bucket}`
    - 完成後會在 s3 bucket 中找到 CodeUri 的物件
2. 發布版本 : `sam publish --template packaged.yml --region us-east-1`
3. 到 AWS SAR, 查看資源資訊
4. 點選部署
5. 查看 ApiGateway, 測試
6. 移除 SAR 
    - 在 aws sar 服務中, 移除資源
    - 在 cloudformation 中, 移除資源
    - 在 s3 移除資源