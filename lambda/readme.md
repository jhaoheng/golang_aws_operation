## 本地測試, 限制
- 一個 go binary 算一個 lambda = 一個 container
- 各個不同的 func example, 參照該 readme.md

## 上傳 aws lambda
1. `zip main.zip ./release/main`
2. 上傳到 lambda func : 注意 handler 名稱 (跟隨 binary's name)