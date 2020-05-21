# readme

- 因為沒有合適的 localhost docker service, 故此為連線到 aws cloud 進行寄信的行為 
- 確定 .aws 有建立, 因為必須使用 aws ses 的線上服務
- 確定 AWS SES 有設定好正確的資訊

# 注意
- 要寄給非驗證過的用戶, 必須將 email 移出 sandbox

# Doc
- SES 設定步驟 : https://docs.aws.amazon.com/zh_tw/ses/latest/DeveloperGuide/verify-email-addresses.html
- golang doc : https://docs.aws.amazon.com/sdk-for-go/api/service/sesv2/
- 程式參考 
    - 範例是 v1 : https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/ses-example-send-email.html
    - 這邊用的是 sesv2
- raw email 參考 : 
    - https://docs.aws.amazon.com/zh_tw/ses/latest/DeveloperGuide/send-email-raw.html
    - 夾帶附加檔案
        - 不支援的附件類型 : https://docs.aws.amazon.com/zh_tw/ses/latest/DeveloperGuide/mime-types-appendix.html
        - 附件 Mine Type 參考 : https://developer.mozilla.org/zh-TW/docs/Web/HTTP/Basics_of_HTTP/MIME_types
            - ex : image/jpeg, text/csv, application/pdf
