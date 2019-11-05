# 關於 session 的運作

## 若是透過 .aws/credential 的方式
- 在 `session.NewSession()` 時, 並不會從 file 中取得參數, 會在執行一次任何與 aws 溝通的行為後, 才會在 Config.Credentials 中得到憑證參數
- 所以如果是 local 端的使用 ... 若不想延宕 request 的速度, 最好在 config 中, 直接指定 Credential 中的參數為 local