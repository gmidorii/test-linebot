

## 追加方法

### curlでの追加
```sh
curl -X POST https://api.line.me/liff/v1/apps \
-H "Authorization: Bearer {channel access token}" \
-H "Content-Type: application/json" \
-d '{
  "view":{
    "type":"full",
    "url":"{}"
  }
}'
```