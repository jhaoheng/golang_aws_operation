## 上傳 aws lambda
1. `GOOS=linux go build`
1. `zip main.zip ./release/main`
2. 上傳到 lambda func : 注意 handler 名稱 (跟隨 binary's name)