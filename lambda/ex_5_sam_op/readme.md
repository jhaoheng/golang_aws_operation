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
## 因以下原因, 必須在本地安裝 awscli & sam
1. 因為 sam 無法指定 aws 憑證資源 (只能設定 profile, 預設讀取同一個路徑)
2. 且 sam 無適當的 docker image, build 資源花太久時間

## 上傳 SAM 到 SAR, 並且從 SAR 中部署資源
> 因有用到 s3 trigger lambda, 故需要設定 s3 bucket policy : https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-template-publishing-applications.html

1. 建立 packaged.yaml : `sam package --template-file template.yaml --output-template-file packaged.yml --s3-bucket {bucket}`
2. 發布版本 : `sam publish --template packaged.yml --region us-east-1`
3. 到 AWS SAR, 查看資源資訊
4. 點選部署
5. 查看 ApiGateway, 測試
6. 移除 SAR 
    - 在 aws sar 服務中, 移除資源
    - 在 cloudformation 中, 移除資源