# 關於 session sdk

1. 使用 aws.Config 可以覆蓋掉 session 中的 `credential` 與相關的環境變數
	- 請參考 : https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
	- credential 可特別指定來源
		- NewStaticCredentials : 靜態變數
		- NewEnvCredentials : 系統環境變數
		- NewSharedCredentials : 讀取 `~/.aws/credentials`
2. 而 session 在預設上(不使用 config), 有以下幾種方法, 可以取得 credential, 優先權由上到下
	1. Environment Variables (優先權最高)
	2. Shared Credentials file (~/.aws/credentials)
	3. Shared Configuration file (if SharedConfig is enabled) : 會讀取 `~/.aws/credential` 中的 
		- region 
		- output
	4. EC2 Instance Metadata (credentials only)
3. 預設 SDK 只會讀取 shared credentials (~/.aws/credentials), 而其他的變數由系統環境變數提供
3. 若在 session 中, 設定 config, 可以覆蓋掉 session 中的參數



# Q
- 若是透過 .aws/credential 的方式
    - 在 `session.NewSession()` 時, 並不會從 file 中取得參數, 會在執行一次任何與 aws 溝通的行為後, 才會在 Config.Credentials 中得到憑證參數
- 要知道可設定哪些環境變數, 可參考
    - https://docs.aws.amazon.com/sdk-for-go/api/aws/session/
- credential
    - golang sdk 中, 當 `SharedConfigEnabled` 時, 會使用 credential 中的 region 參數
    - 若不使用 credential 中的 region 參數, 則預設使用環境變數或由 aws.Config 進行設定