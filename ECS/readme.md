- https://docs.aws.amazon.com/zh_tw/AmazonECS/latest/developerguide/task-metadata-endpoint-v3.html
- https://noise.getoto.net/2019/03/27/a-guide-to-locally-testing-containers-with-amazon-ecs-local-endpoints-and-docker-compose/



- ECS_CONTAINER_METADATA_URI: "http://169.254.170.2/v3"
- This is the environment variable defined by the V3 metadata spec.


這邊的 .aws 需要取得正確的 credentials, 否則 ecs-local-endpoints 不會發 /cred 給其他的 container 使用


所以只要塞到 `ecs-local-endpoints` 的 credentials 是有授權的
透過 `ecs-local-endpoints` 拿到的 iam 都可以 access 到相關的服務中

- 查詢 container 登入資料 : `curl 169.254.170.2$AWS_CONTAINER_CREDENTIALS_RELATIVE_URI`
    - https://docs.aws.amazon.com/zh_tw/AmazonECS/latest/developerguide/task-iam-roles.html
- 輸出
```
{
    "AccessKeyId": "ACCESS_KEY_ID",
    "Expiration": "EXPIRATION_DATE",
    "RoleArn": "TASK_ROLE_ARN",
    "SecretAccessKey": "SECRET_ACCESS_KEY",
    "Token": "SECURITY_TOKEN_STRING"
}
```

# 任務中繼資料端點第 3 版路徑 ($ECS_CONTAINER_METADATA_URI)
> https://docs.aws.amazon.com/zh_tw/AmazonECS/latest/developerguide/task-metadata-endpoint-v3.html

- ${ECS_CONTAINER_METADATA_URI}
- ${ECS_CONTAINER_METADATA_URI}/task : 啟動的 task 相關資訊, 可以存起來
- ${ECS_CONTAINER_METADATA_URI}/stats
- ${ECS_CONTAINER_METADATA_URI}/task/stats
